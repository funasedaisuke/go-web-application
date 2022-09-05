package handler

import (
	"net/http"
	"testing"

    "github.com/funasedaisuke/go-web-application/store"
	"github.com/funasedaisuke/go-web-application/entity"
	"github.com/funasedaisuke/go-web-application/testutil"
)

func TestAddTask(t *testing.T){
	t.Parallel()
	type want struct {
		status int
		rspFile string
	}
	tests := map[string]struct{
		reqFile string
        want want
	}{
		"ok":{
			reqFile: "testdata/add_task/ok_req.json.golden",
        	want: want{
				status: http.StatusOK,
		        rspFile: "testdata/add_task/ok_req_rsp.json.golden",
			},
		},
		"badRequest":{
			reqFile: "testdata/add_task/bad_req.json.golden",
        	want: want{
				status: http.StatusBadRequest,
		        rspFile: "testdata/add_task/bad_req_rsp.json.golden",
			},
		},		
	}

	for n,tt := range tests{
		tt := tt
		t.Run(n,func(t *testing.T)){
			t.Parallel()
			w := httptest.NewRecoder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/tasks",
				bytes.NewRecoder(testutil.Loadfile(t,tt.readfile)),

			)
			sut := AddTask{
				Store: &store.TaskStore{
					Tasks: map[entity.TaskID]*entity.Task{},
				}
				validator: validator.New()
			}
			sut.ServerHTTP(w,r)

			resp :=w.Result()
			testutil.AssertResponse(t,
			resp, tt.want.status,testutil.LoadFile())
		}


	}

	
}