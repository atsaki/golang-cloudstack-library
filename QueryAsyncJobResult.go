package cloudstack

import (
	"encoding/json"
	"errors"
	"fmt"
)

type QueryAsyncJobResultParameter struct {
	JobId NullString
}

type AsyncJobResult struct {
	AccountId       NullString      `json:"accountid"`
	Cmd             NullString      `json:"cmd"`
	Created         NullString      `json:"created"`
	JobId           NullString      `json:"jobid"`
	JobInstanceId   NullString      `json:"jobinstanceid"`
	JobInstanceType NullString      `json:"jobinstancetype"`
	JobProcsSatus   NullNumber      `json:"jobprocstatus"`
	JobResult       json.RawMessage `json:"jobresult"`
	JobResultCode   NullNumber      `json:"jobresultcode"`
	JobResultType   NullString      `json:"jobresulttype"`
	JobStatus       NullNumber      `json:"jobstatus"`
	UserId          NullString      `json:"userid"`
}

func (c *Client) QueryAsyncJobResult(p *QueryAsyncJobResultParameter) (job *AsyncJobResult, err error) {

	job = new(AsyncJobResult)

	b, err := c.Request("queryAsyncJobResult", convertParamToMap(p))
	if err != nil {
		return job, err
	}

	err = json.Unmarshal(b, job)
	if err != nil {
		return job, fmt.Errorf("Failed to unmarshal (%s): %s", err, string(b))
	}

	if !job.JobStatus.IsNil() &&
		(job.JobStatus.String() == "0" || job.JobStatus.String() == "1") {
		return job, nil
	}

	errortext := getErrorText(b)
	if errortext != "" {
		return job, errors.New(errortext)
	}
	return job, fmt.Errorf("QueryAsyncJobResult didn't finished correctly : jobstatus: %v", job.JobStatus)
}
