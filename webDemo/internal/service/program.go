package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	//"fmt"
	"io/ioutil"
	"webDemo/internal/model"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type CreateProgramRequest struct {
	Program_id   uint32 `form:"program_id"`
	Program_name string `form:"program_name" binding:"max=100"`
	Content      string `form:"content" binding:"required,min=10,max=500"`
	Ptype        string `form:"ptype" binding:"required,min=2,max=50"`
	Answer       string `form:"answer" binding:"required,min=1,max=100"`
	Difficulty   string `form:"difficulty" binding:"required,min=1,max=50"`
}

func (svc *Service) CreateProgram(param *CreateProgramRequest) error {
	return svc.dao.CreateProgram(param.Program_name, param.Content, param.Ptype, param.Answer, param.Difficulty)
}

type DeleteProgramRequest struct {
	Program_id uint32 `form:"program_id"`
}

func (svc *Service) DeleteProgram(program_id uint32) error {
	return svc.dao.DeleteProgram(program_id)
}

type SubmitProgramRequest struct {
	User_id    uint32 `form:"user_id"`
	Program_id uint32 `form:"program_id"`
	Answer     string `form:"answer" binding:"required,min=1,max=50000"`
}

type SubmitDockerProgramRequest struct {
	User_id    uint32 `form:"user_id"`
	AnswerCode string `form:"AnswerCode" binding:"required,min=1,max=50000"`
}

type ReturnProgramListRequest struct {
	Program_id   uint32 `form:"program_id"`
	Program_name string `form:"program_name" binding:"max=100"`
	Answer       string `form:"answer" binding:"required,min=1,max=100"`
	Difficulty   string `form:"difficulty" binding:"required,min=1,max=50"`
}

func (svc *Service) ReturnProgramList(param *ReturnProgramListRequest) ([]model.Program, error) {
	return svc.dao.ReturnProgramList(param.Program_id, param.Program_name, param.Answer, param.Difficulty)
}

type ReturnProgramDetailRequest struct {
	Program_id   uint32 `form:"program_id"`
	Program_name string `form:"program_name" binding:"max=100"`
}

func (svc *Service) ReturnProgramDetail(param *ReturnProgramDetailRequest) ([]model.Program, error) {
	return svc.dao.ReturnProgramDetail(param.Program_id, param.Program_name)
}

func (svc *Service) SubmitProgram(param *SubmitProgramRequest) (bool, error) {
	program, err := svc.dao.ReturnProgramDetail(param.Program_id, "")
	if len(program) < 1 {
		return false, errors.New("no program")
	}
	if err != nil {
		return false, err
	}
	if program[0].Answer == param.Answer {
		err = svc.dao.CreateRecord(param.Program_id, param.User_id, "pass",program[0].Difficulty)
		if err != nil {
			return false, err
		}
		return true, nil
	} else {
		err = svc.dao.CreateRecord(param.Program_id, param.User_id, "not pass",program[0].Difficulty)
		if err != nil {
			return false, err
		}
		return false, nil
	}
}

type ReturnRecord struct {
	Record_id  uint32 `form:"record_id"`
	User_id    uint32 `form:"user_id"`
	Program_id uint32 `form:"program_id"`
	Status     string `form:"status"`
	Difficulty string `form:"difficulty"`
}

func (svc *Service) ReturnRecord(param *ReturnRecord) ([]model.Record, error) {
	fmt.Printf("service id = %d", param.User_id)
	return svc.dao.ReturnRecord(param.Record_id, param.Program_id, param.User_id, param.Status, param.Difficulty)
}

// return 1 for compile success, 0 for time out
func startDocker(program_name string) (int, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		return 0, err
	}
	env := "PROGRAM=" + program_name
	fmt.Printf("docker files:")
	binds_arg := []string{"/home/yeye/program/webDemo/user_submit:/go/volume"}
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Env:   []string{env},
		Image: "web:5",
		Cmd:   []string{"/go/check.sh"},
		//Cmd: []string{"ls", "/go/volume"},
	},
		&container.HostConfig{
			Binds: binds_arg,
			//VolumeDriver: "/home/yeye/program/webDemo/user_submit",
		}, nil, nil, "")
	if err != nil {
		return 0, err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return 0, err
	}
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return 0, err
		}
	case <-statusCh:
		return 1, nil
	case <-time.NewTimer(time.Second * 5).C:
		return 0, nil
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return 0, err
	}
	io.Copy(os.Stdout, out)
	return 0, err
}

// 1=success, 2=compile err, 3 = wrong answer, 4=has err, 5=time out
func (svc *Service) SubmitDockerProgram(param *SubmitDockerProgramRequest, program_id int) (int, string, error) {
	program, err := svc.dao.ReturnProgramDetail(uint32(program_id), "")
	if len(program) < 1 {
		return 4, "", errors.New("no program")
	}
	if err != nil {
		return 4, "", err
	}

	// 创建programname = userid_programid_time+.go, 创建代码文件
	unixTime := time.Now().Unix()
	programName := strconv.Itoa(int(param.User_id)) + "_" + strconv.Itoa(program_id) + "_" + strconv.FormatInt(unixTime, 10) + ".go"
	fmt.Printf("programid = %v\n", programName)

	programFile, error := os.OpenFile("./user_submit/"+programName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Printf("%v: cannot create file", error)
		return 4, "", err
	}
	_, err = programFile.WriteString(param.AnswerCode)
	if err != nil {
		fmt.Printf("%v: cannot write", error)
		return 4, "", err
	}
	programFile.Close()

	ret, err := startDocker(programName)
	if err != nil {
		return 4, "", err
	} else if ret == 0 {
		return 5, "time out of 10s", nil
	}

	//不知未存在会不会爆error
	outputFile, err := ioutil.ReadFile("./user_submit/" + programName + "-out")
	if err != nil {
		fmt.Printf("%v: output file not exit", error)
		return 4, "", err
	}

	fmt.Printf("output data = %v, real answer = %v", string(outputFile), program[0].Answer)

	if program[0].Answer == string(outputFile) {
		err = svc.dao.CreateRecord(uint32(program_id), param.User_id, "pass",program[0].Difficulty)
		if err != nil {
			return 4, "", err
		}
		return 1, "", nil
	} else {
		// 不相同判断是否是编译错误
		svc.dao.CreateRecord(uint32(program_id), param.User_id, "not pass",program[0].Difficulty)
		errFile, err := ioutil.ReadFile("./user_submit/" + programName + "-err")
		if err != nil {
			fmt.Printf("%v: err file not exit", error)
		}
		if string(errFile) == "" {
			return 3, string(outputFile), err
		} else {
			return 2, string(errFile), err
		}
	}
}
