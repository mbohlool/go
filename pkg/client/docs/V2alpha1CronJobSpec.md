# V2alpha1CronJobSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConcurrencyPolicy** | **string** | Specifies how to treat concurrent executions of a Job. Defaults to Allow. | [optional] [default to null]
**FailedJobsHistoryLimit** | **int32** | The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified. | [optional] [default to null]
**JobTemplate** | [**V2alpha1JobTemplateSpec**](v2alpha1.JobTemplateSpec.md) | Specifies the job that will be created when executing a CronJob. | [default to null]
**Schedule** | **string** | The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron. | [default to null]
**StartingDeadlineSeconds** | **int64** | Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones. | [optional] [default to null]
**SuccessfulJobsHistoryLimit** | **int32** | The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified. | [optional] [default to null]
**Suspend** | **bool** | This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


