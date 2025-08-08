# Dify API Documentation

**Generated at:** 2025-08-05T04:40:31.801024
**Total Blueprints:** 6
**Total Endpoints:** 470

## Table of Contents
- [service_api](#service-api)
- [web](#web)
- [console](#console)
- [files](#files)
- [inner_api](#inner-api)
- [mcp](#mcp)

## service_api
**URL Prefix:** `/v1`

| Method | URL | Function | file |
|--------|-----|----------|-------------|
| GET | `/v1/` | `indexapi` | controllers.service_api.index.py |
| POST | `/v1/apps/annotation-reply/<string:action>` | `annotationreplyactionapi` | controllers.service_api.app.annotation.py |
| GET | `/v1/apps/annotation-reply/<string:action>/status/<uuid:job_id>` | `annotationreplyactionstatusapi` | controllers.service_api.app.annotation.py |
| POST, GET | `/v1/apps/annotations` | `annotationlistapi` | controllers.service_api.app.annotation.py |
| PUT, DELETE | `/v1/apps/annotations/<uuid:annotation_id>` | `annotationupdatedeleteapi` | controllers.service_api.app.annotation.py |
| GET | `/v1/parameters` | `appparameterapi` | controllers.service_api.app.app.py |
| GET | `/v1/meta` | `appmetaapi` | controllers.service_api.app.app.py |
| GET | `/v1/info` | `appinfoapi` | controllers.service_api.app.app.py |
| POST | `/v1/audio-to-text` | `audioapi` | controllers.service_api.app.audio.py |
| POST | `/v1/text-to-audio` | `textapi` | controllers.service_api.app.audio.py |
| POST | `/v1/completion-messages` | `completionapi` | controllers.service_api.app.completion.py |
| POST | `/v1/completion-messages/<string:task_id>/stop` | `completionstopapi` | controllers.service_api.app.completion.py |
| POST | `/v1/chat-messages` | `chatapi` | controllers.service_api.app.completion.py |
| POST | `/v1/chat-messages/<string:task_id>/stop` | `chatstopapi` | controllers.service_api.app.completion.py |
| POST | `/v1/conversations/<uuid:c_id>/name` | `conversation_name` | controllers.service_api.app.conversation.py |
| GET | `/v1/conversations` | `conversationapi` | controllers.service_api.app.conversation.py |
| DELETE | `/v1/conversations/<uuid:c_id>` | `conversation_detail` | controllers.service_api.app.conversation.py |
| GET | `/v1/conversations/<uuid:c_id>/variables` | `conversation_variables` | controllers.service_api.app.conversation.py |
| POST | `/v1/files/upload` | `fileapi` | controllers.service_api.app.file.py |
| GET | `/v1/messages` | `messagelistapi` | controllers.service_api.app.message.py |
| POST | `/v1/messages/<uuid:message_id>/feedbacks` | `messagefeedbackapi` | controllers.service_api.app.message.py |
| GET | `/v1/messages/<uuid:message_id>/suggested` | `messagesuggestedapi` | controllers.service_api.app.message.py |
| GET | `/v1/app/feedbacks` | `appgetfeedbacksapi` | controllers.service_api.app.message.py |
| GET | `/v1/site` | `appsiteapi` | controllers.service_api.app.site.py |
| POST | `/v1/workflows/run` | `workflowrunapi` | controllers.service_api.app.workflow.py |
| GET | `/v1/workflows/run/<string:workflow_run_id>` | `workflowrundetailapi` | controllers.service_api.app.workflow.py |
| POST | `/v1/workflows/tasks/<string:task_id>/stop` | `workflowtaskstopapi` | controllers.service_api.app.workflow.py |
| GET | `/v1/workflows/logs` | `workflowapplogapi` | controllers.service_api.app.workflow.py |
| POST, GET | `/v1/datasets` | `datasetlistapi` | controllers.service_api.dataset.dataset.py |
| PATCH, GET, DELETE | `/v1/datasets/<uuid:dataset_id>` | `datasetapi` | controllers.service_api.dataset.dataset.py |
| PATCH | `/v1/datasets/<uuid:dataset_id>/documents/status/<string:action>` | `documentstatusapi` | controllers.service_api.dataset.dataset.py |
| PATCH, POST, GET, DELETE | `/v1/datasets/tags` | `datasettagsapi` | controllers.service_api.dataset.dataset.py |
| POST | `/v1/datasets/tags/binding` | `datasettagbindingapi` | controllers.service_api.dataset.dataset.py |
| POST | `/v1/datasets/tags/unbinding` | `datasettagunbindingapi` | controllers.service_api.dataset.dataset.py |
| GET | `/v1/datasets/<uuid:dataset_id>/tags` | `datasettagsbindingstatusapi` | controllers.service_api.dataset.dataset.py |
| POST | `/v1/datasets/<uuid:dataset_id>/document/create_by_text` | `documentaddbytextapi` | controllers.service_api.dataset.document.py |
| POST | `/v1/datasets/<uuid:dataset_id>/document/create-by-text` | `documentaddbytextapi` | controllers.service_api.dataset.document.py |
| POST | `/v1/datasets/<uuid:dataset_id>/document/create_by_file` | `documentaddbyfileapi` | controllers.service_api.dataset.document.py |
| POST | `/v1/datasets/<uuid:dataset_id>/document/create-by-file` | `documentaddbyfileapi` | controllers.service_api.dataset.document.py |
| POST | `/v1/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/update_by_text` | `documentupdatebytextapi` | controllers.service_api.dataset.document.py |
| POST | `/v1/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/update-by-text` | `documentupdatebytextapi` | controllers.service_api.dataset.document.py |
| POST | `/v1/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/update_by_file` | `documentupdatebyfileapi` | controllers.service_api.dataset.document.py |
| POST | `/v1/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/update-by-file` | `documentupdatebyfileapi` | controllers.service_api.dataset.document.py |
| DELETE | `/v1/datasets/<uuid:dataset_id>/documents/<uuid:document_id>` | `documentdeleteapi` | controllers.service_api.dataset.document.py |
| GET | `/v1/datasets/<uuid:dataset_id>/documents` | `documentlistapi` | controllers.service_api.dataset.document.py |
| GET | `/v1/datasets/<uuid:dataset_id>/documents/<string:batch>/indexing-status` | `documentindexingstatusapi` | controllers.service_api.dataset.document.py |
| GET | `/v1/datasets/<uuid:dataset_id>/documents/<uuid:document_id>` | `documentdetailapi` | controllers.service_api.dataset.document.py |
| POST | `/v1/datasets/<uuid:dataset_id>/hit-testing` | `hittestingapi` | controllers.service_api.dataset.hit_testing.py |
| POST | `/v1/datasets/<uuid:dataset_id>/retrieve` | `hittestingapi` | controllers.service_api.dataset.hit_testing.py |
| POST, GET | `/v1/datasets/<uuid:dataset_id>/metadata` | `datasetmetadatacreateserviceapi` | controllers.service_api.dataset.metadata.py |
| PATCH, DELETE | `/v1/datasets/<uuid:dataset_id>/metadata/<uuid:metadata_id>` | `datasetmetadataserviceapi` | controllers.service_api.dataset.metadata.py |
| GET | `/v1/datasets/metadata/built-in` | `datasetmetadatabuiltinfieldserviceapi` | controllers.service_api.dataset.metadata.py |
| POST | `/v1/datasets/<uuid:dataset_id>/metadata/built-in/<string:action>` | `datasetmetadatabuiltinfieldactionserviceapi` | controllers.service_api.dataset.metadata.py |
| POST | `/v1/datasets/<uuid:dataset_id>/documents/metadata` | `documentmetadataeditserviceapi` | controllers.service_api.dataset.metadata.py |
| POST, GET | `/v1/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/segments` | `segmentapi` | controllers.service_api.dataset.segment.py |
| POST, GET, DELETE | `/v1/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/segments/<uuid:segment_id>` | `datasetsegmentapi` | controllers.service_api.dataset.segment.py |
| POST, GET | `/v1/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/segments/<uuid:segment_id>/child_chunks` | `childchunkapi` | controllers.service_api.dataset.segment.py |
| PATCH, DELETE | `/v1/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/segments/<uuid:segment_id>/child_chunks/<uuid:child_chunk_id>` | `datasetchildchunkapi` | controllers.service_api.dataset.segment.py |
| GET | `/v1/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/upload-file` | `uploadfileapi` | controllers.service_api.dataset.upload_file.py |
| GET | `/v1/workspaces/current/models/model-types/<string:model_type>` | `modelprovideravailablemodelapi` | controllers.service_api.workspace.models.py |

## web
**URL Prefix:** `/api`

| Method | URL | Function | file |
|--------|-----|----------|-------------|
| POST | `/api/files/upload` | `fileapi` | controllers.web.files.py |
| GET | `/api/remote-files/<path:url>` | `remotefileinfoapi` | controllers.web.remote_files.py |
| POST | `/api/remote-files/upload` | `remotefileuploadapi` | controllers.web.remote_files.py |
| GET | `/api/parameters` | `appparameterapi` | controllers.web.app.py |
| GET | `/api/meta` | `appmeta` | controllers.web.app.py |
| GET | `/api/webapp/access-mode` | `appaccessmode` | controllers.web.app.py |
| GET | `/api/webapp/permission` | `appwebauthpermission` | controllers.web.app.py |
| POST | `/api/audio-to-text` | `audioapi` | controllers.web.audio.py |
| POST | `/api/text-to-audio` | `textapi` | controllers.web.audio.py |
| POST | `/api/completion-messages` | `completionapi` | controllers.web.completion.py |
| POST | `/api/completion-messages/<string:task_id>/stop` | `completionstopapi` | controllers.web.completion.py |
| POST | `/api/chat-messages` | `chatapi` | controllers.web.completion.py |
| POST | `/api/chat-messages/<string:task_id>/stop` | `chatstopapi` | controllers.web.completion.py |
| POST | `/api/conversations/<uuid:c_id>/name` | `web_conversation_name` | controllers.web.conversation.py |
| GET | `/api/conversations` | `conversationlistapi` | controllers.web.conversation.py |
| DELETE | `/api/conversations/<uuid:c_id>` | `conversationapi` | controllers.web.conversation.py |
| PATCH | `/api/conversations/<uuid:c_id>/pin` | `conversationpinapi` | controllers.web.conversation.py |
| PATCH | `/api/conversations/<uuid:c_id>/unpin` | `conversationunpinapi` | controllers.web.conversation.py |
| GET | `/api/system-features` | `systemfeatureapi` | controllers.web.feature.py |
| POST | `/api/forgot-password` | `forgotpasswordsendemailapi` | controllers.web.forgot_password.py |
| POST | `/api/forgot-password/validity` | `forgotpasswordcheckapi` | controllers.web.forgot_password.py |
| POST | `/api/forgot-password/resets` | `forgotpasswordresetapi` | controllers.web.forgot_password.py |
| POST | `/api/login` | `loginapi` | controllers.web.login.py |
| POST | `/api/email-code-login` | `emailcodeloginsendemailapi` | controllers.web.login.py |
| POST | `/api/email-code-login/validity` | `emailcodeloginapi` | controllers.web.login.py |
| GET | `/api/messages` | `messagelistapi` | controllers.web.message.py |
| POST | `/api/messages/<uuid:message_id>/feedbacks` | `messagefeedbackapi` | controllers.web.message.py |
| GET | `/api/messages/<uuid:message_id>/more-like-this` | `messagemorelikethisapi` | controllers.web.message.py |
| GET | `/api/messages/<uuid:message_id>/suggested-questions` | `messagesuggestedquestionapi` | controllers.web.message.py |
| GET | `/api/passport` | `passportresource` | controllers.web.passport.py |
| POST, GET | `/api/saved-messages` | `savedmessagelistapi` | controllers.web.saved_message.py |
| DELETE | `/api/saved-messages/<uuid:message_id>` | `savedmessageapi` | controllers.web.saved_message.py |
| GET | `/api/site` | `appsiteapi` | controllers.web.site.py |
| POST | `/api/workflows/run` | `workflowrunapi` | controllers.web.workflow.py |
| POST | `/api/workflows/tasks/<string:task_id>/stop` | `workflowtaskstopapi` | controllers.web.workflow.py |

## console
**URL Prefix:** `/console/api`

| Method | URL | Function | file |
|--------|-----|----------|-------------|
| POST, GET | `/console/api/files/upload` | `fileapi` | controllers.console.files.py |
| GET | `/console/api/files/<uuid:file_id>/preview` | `filepreviewapi` | controllers.console.files.py |
| GET | `/console/api/files/support-type` | `filesupporttypeapi` | controllers.console.files.py |
| GET | `/console/api/remote-files/<path:url>` | `remotefileinfoapi` | controllers.console.remote_files.py |
| POST | `/console/api/remote-files/upload` | `remotefileuploadapi` | controllers.console.remote_files.py |
| POST | `/console/api/apps/imports` | `appimportapi` | controllers.console.app.app_import.py |
| POST | `/console/api/apps/imports/<string:import_id>/confirm` | `appimportconfirmapi` | controllers.console.app.app_import.py |
| GET | `/console/api/apps/imports/<string:app_id>/check-dependencies` | `appimportcheckdependenciesapi` | controllers.console.app.app_import.py |
| POST | `/console/api/admin/insert-explore-apps` | `insertexploreapplistapi` | controllers.console.admin.py |
| DELETE | `/console/api/admin/insert-explore-apps/<uuid:app_id>` | `insertexploreappapi` | controllers.console.admin.py |
| POST, GET | `/console/api/apps/<uuid:resource_id>/api-keys` | `appapikeylistresource` | controllers.console.apikey.py |
| DELETE | `/console/api/apps/<uuid:resource_id>/api-keys/<uuid:api_key_id>` | `appapikeyresource` | controllers.console.apikey.py |
| POST, GET | `/console/api/datasets/<uuid:resource_id>/api-keys` | `datasetapikeylistresource` | controllers.console.apikey.py |
| DELETE | `/console/api/datasets/<uuid:resource_id>/api-keys/<uuid:api_key_id>` | `datasetapikeyresource` | controllers.console.apikey.py |
| GET | `/console/api/code-based-extension` | `codebasedextensionapi` | controllers.console.extension.py |
| POST, GET | `/console/api/api-based-extension` | `apibasedextensionapi` | controllers.console.extension.py |
| POST, GET, DELETE | `/console/api/api-based-extension/<uuid:id>` | `apibasedextensiondetailapi` | controllers.console.extension.py |
| GET | `/console/api/features` | `featureapi` | controllers.console.feature.py |
| GET | `/console/api/system-features` | `systemfeatureapi` | controllers.console.feature.py |
| GET | `/console/api/ping` | `pingapi` | controllers.console.ping.py |
| POST, GET | `/console/api/init` | `initvalidateapi` | controllers.console.init_validate.py |
| POST, GET | `/console/api/setup` | `setupapi` | controllers.console.setup.py |
| GET | `/console/api/version` | `versionapi` | controllers.console.version.py |
| GET | `/console/api/app/prompt-templates` | `advancedprompttemplatelist` | controllers.console.app.advanced_prompt_template.py |
| GET | `/console/api/apps/<uuid:app_id>/agent/logs` | `agentlogapi` | controllers.console.app.agent.py |
| POST | `/console/api/apps/<uuid:app_id>/annotation-reply/<string:action>` | `annotationreplyactionapi` | controllers.console.app.annotation.py |
| GET | `/console/api/apps/<uuid:app_id>/annotation-reply/<string:action>/status/<uuid:job_id>` | `annotationreplyactionstatusapi` | controllers.console.app.annotation.py |
| GET, DELETE | `/console/api/apps/<uuid:app_id>/annotations` | `annotationlistapi` | controllers.console.app.annotation.py |
| GET | `/console/api/apps/<uuid:app_id>/annotations/export` | `annotationexportapi` | controllers.console.app.annotation.py |
| POST, DELETE | `/console/api/apps/<uuid:app_id>/annotations/<uuid:annotation_id>` | `annotationupdatedeleteapi` | controllers.console.app.annotation.py |
| POST | `/console/api/apps/<uuid:app_id>/annotations/batch-import` | `annotationbatchimportapi` | controllers.console.app.annotation.py |
| GET | `/console/api/apps/<uuid:app_id>/annotations/batch-import-status/<uuid:job_id>` | `annotationbatchimportstatusapi` | controllers.console.app.annotation.py |
| GET | `/console/api/apps/<uuid:app_id>/annotations/<uuid:annotation_id>/hit-histories` | `annotationhithistorylistapi` | controllers.console.app.annotation.py |
| GET | `/console/api/apps/<uuid:app_id>/annotation-setting` | `appannotationsettingdetailapi` | controllers.console.app.annotation.py |
| POST | `/console/api/apps/<uuid:app_id>/annotation-settings/<uuid:annotation_setting_id>` | `appannotationsettingupdateapi` | controllers.console.app.annotation.py |
| POST, GET | `/console/api/apps` | `applistapi` | controllers.console.app.app.py |
| PUT, GET, DELETE | `/console/api/apps/<uuid:app_id>` | `appapi` | controllers.console.app.app.py |
| POST | `/console/api/apps/<uuid:app_id>/copy` | `appcopyapi` | controllers.console.app.app.py |
| GET | `/console/api/apps/<uuid:app_id>/export` | `appexportapi` | controllers.console.app.app.py |
| POST | `/console/api/apps/<uuid:app_id>/name` | `appnameapi` | controllers.console.app.app.py |
| POST | `/console/api/apps/<uuid:app_id>/icon` | `appiconapi` | controllers.console.app.app.py |
| POST | `/console/api/apps/<uuid:app_id>/site-enable` | `appsitestatus` | controllers.console.app.app.py |
| POST | `/console/api/apps/<uuid:app_id>/api-enable` | `appapistatus` | controllers.console.app.app.py |
| POST, GET | `/console/api/apps/<uuid:app_id>/trace` | `apptraceapi` | controllers.console.app.app.py |
| POST | `/console/api/apps/<uuid:app_id>/audio-to-text` | `chatmessageaudioapi` | controllers.console.app.audio.py |
| POST | `/console/api/apps/<uuid:app_id>/text-to-audio` | `chatmessagetextapi` | controllers.console.app.audio.py |
| GET | `/console/api/apps/<uuid:app_id>/text-to-audio/voices` | `textmodesapi` | controllers.console.app.audio.py |
| POST | `/console/api/apps/<uuid:app_id>/completion-messages` | `completionmessageapi` | controllers.console.app.completion.py |
| POST | `/console/api/apps/<uuid:app_id>/completion-messages/<string:task_id>/stop` | `completionmessagestopapi` | controllers.console.app.completion.py |
| POST | `/console/api/apps/<uuid:app_id>/chat-messages` | `chatmessageapi` | controllers.console.app.completion.py |
| POST | `/console/api/apps/<uuid:app_id>/chat-messages/<string:task_id>/stop` | `chatmessagestopapi` | controllers.console.app.completion.py |
| GET | `/console/api/apps/<uuid:app_id>/completion-conversations` | `completionconversationapi` | controllers.console.app.conversation.py |
| GET, DELETE | `/console/api/apps/<uuid:app_id>/completion-conversations/<uuid:conversation_id>` | `completionconversationdetailapi` | controllers.console.app.conversation.py |
| GET | `/console/api/apps/<uuid:app_id>/chat-conversations` | `chatconversationapi` | controllers.console.app.conversation.py |
| GET, DELETE | `/console/api/apps/<uuid:app_id>/chat-conversations/<uuid:conversation_id>` | `chatconversationdetailapi` | controllers.console.app.conversation.py |
| GET | `/console/api/apps/<uuid:app_id>/conversation-variables` | `conversationvariablesapi` | controllers.console.app.conversation_variables.py |
| POST | `/console/api/rule-generate` | `rulegenerateapi` | controllers.console.app.generator.py |
| POST | `/console/api/rule-code-generate` | `rulecodegenerateapi` | controllers.console.app.generator.py |
| POST | `/console/api/rule-structured-output-generate` | `rulestructuredoutputgenerateapi` | controllers.console.app.generator.py |
| POST, PUT, GET | `/console/api/apps/<uuid:app_id>/server` | `appmcpservercontroller` | controllers.console.app.mcp_server.py |
| GET | `/console/api/apps/<uuid:server_id>/server/refresh` | `appmcpserverrefreshcontroller` | controllers.console.app.mcp_server.py |
| GET | `/console/api/apps/<uuid:app_id>/chat-messages/<uuid:message_id>/suggested-questions` | `messagesuggestedquestionapi` | controllers.console.app.message.py |
| GET | `/console/api/apps/<uuid:app_id>/chat-messages` | `console_chat_messages` | controllers.console.app.message.py |
| POST | `/console/api/apps/<uuid:app_id>/feedbacks` | `messagefeedbackapi` | controllers.console.app.message.py |
| POST | `/console/api/apps/<uuid:app_id>/annotations` | `messageannotationapi` | controllers.console.app.message.py |
| GET | `/console/api/apps/<uuid:app_id>/annotations/count` | `messageannotationcountapi` | controllers.console.app.message.py |
| GET | `/console/api/apps/<uuid:app_id>/messages/<uuid:message_id>` | `console_message` | controllers.console.app.message.py |
| POST | `/console/api/apps/<uuid:app_id>/model-config` | `modelconfigresource` | controllers.console.app.model_config.py |
| PATCH, POST, GET, DELETE | `/console/api/apps/<uuid:app_id>/trace-config` | `traceappconfigapi` | controllers.console.app.ops_trace.py |
| POST | `/console/api/apps/<uuid:app_id>/site` | `appsite` | controllers.console.app.site.py |
| POST | `/console/api/apps/<uuid:app_id>/site/access-token-reset` | `appsiteaccesstokenreset` | controllers.console.app.site.py |
| GET | `/console/api/apps/<uuid:app_id>/statistics/daily-messages` | `dailymessagestatistic` | controllers.console.app.statistic.py |
| GET | `/console/api/apps/<uuid:app_id>/statistics/daily-conversations` | `dailyconversationstatistic` | controllers.console.app.statistic.py |
| GET | `/console/api/apps/<uuid:app_id>/statistics/daily-end-users` | `dailyterminalsstatistic` | controllers.console.app.statistic.py |
| GET | `/console/api/apps/<uuid:app_id>/statistics/token-costs` | `dailytokencoststatistic` | controllers.console.app.statistic.py |
| GET | `/console/api/apps/<uuid:app_id>/statistics/average-session-interactions` | `averagesessioninteractionstatistic` | controllers.console.app.statistic.py |
| GET | `/console/api/apps/<uuid:app_id>/statistics/user-satisfaction-rate` | `usersatisfactionratestatistic` | controllers.console.app.statistic.py |
| GET | `/console/api/apps/<uuid:app_id>/statistics/average-response-time` | `averageresponsetimestatistic` | controllers.console.app.statistic.py |
| GET | `/console/api/apps/<uuid:app_id>/statistics/tokens-per-second` | `tokenspersecondstatistic` | controllers.console.app.statistic.py |
| POST, GET | `/console/api/apps/<uuid:app_id>/workflows/draft` | `draftworkflowapi` | controllers.console.app.workflow.py |
| GET | `/console/api/apps/<uuid:app_id>/workflows/draft/config` | `workflowconfigapi` | controllers.console.app.workflow.py |
| POST | `/console/api/apps/<uuid:app_id>/advanced-chat/workflows/draft/run` | `advancedchatdraftworkflowrunapi` | controllers.console.app.workflow.py |
| POST | `/console/api/apps/<uuid:app_id>/workflows/draft/run` | `draftworkflowrunapi` | controllers.console.app.workflow.py |
| POST | `/console/api/apps/<uuid:app_id>/workflow-runs/tasks/<string:task_id>/stop` | `workflowtaskstopapi` | controllers.console.app.workflow.py |
| POST | `/console/api/apps/<uuid:app_id>/workflows/draft/nodes/<string:node_id>/run` | `draftworkflownoderunapi` | controllers.console.app.workflow.py |
| POST | `/console/api/apps/<uuid:app_id>/advanced-chat/workflows/draft/iteration/nodes/<string:node_id>/run` | `advancedchatdraftruniterationnodeapi` | controllers.console.app.workflow.py |
| POST | `/console/api/apps/<uuid:app_id>/workflows/draft/iteration/nodes/<string:node_id>/run` | `workflowdraftruniterationnodeapi` | controllers.console.app.workflow.py |
| POST | `/console/api/apps/<uuid:app_id>/advanced-chat/workflows/draft/loop/nodes/<string:node_id>/run` | `advancedchatdraftrunloopnodeapi` | controllers.console.app.workflow.py |
| POST | `/console/api/apps/<uuid:app_id>/workflows/draft/loop/nodes/<string:node_id>/run` | `workflowdraftrunloopnodeapi` | controllers.console.app.workflow.py |
| POST, GET | `/console/api/apps/<uuid:app_id>/workflows/publish` | `publishedworkflowapi` | controllers.console.app.workflow.py |
| GET | `/console/api/apps/<uuid:app_id>/workflows` | `publishedallworkflowapi` | controllers.console.app.workflow.py |
| GET | `/console/api/apps/<uuid:app_id>/workflows/default-workflow-block-configs` | `defaultblockconfigsapi` | controllers.console.app.workflow.py |
| GET | `/console/api/apps/<uuid:app_id>/workflows/default-workflow-block-configs/<string:block_type>` | `defaultblockconfigapi` | controllers.console.app.workflow.py |
| POST | `/console/api/apps/<uuid:app_id>/convert-to-workflow` | `converttoworkflowapi` | controllers.console.app.workflow.py |
| PATCH, DELETE | `/console/api/apps/<uuid:app_id>/workflows/<string:workflow_id>` | `workflowbyidapi` | controllers.console.app.workflow.py |
| GET | `/console/api/apps/<uuid:app_id>/workflows/draft/nodes/<string:node_id>/last-run` | `draftworkflownodelastrunapi` | controllers.console.app.workflow.py |
| GET | `/console/api/apps/<uuid:app_id>/workflow-app-logs` | `workflowapplogapi` | controllers.console.app.workflow_app_log.py |
| GET, DELETE | `/console/api/apps/<uuid:app_id>/workflows/draft/variables` | `workflowvariablecollectionapi` | controllers.console.app.workflow_draft_variable.py |
| GET, DELETE | `/console/api/apps/<uuid:app_id>/workflows/draft/nodes/<string:node_id>/variables` | `nodevariablecollectionapi` | controllers.console.app.workflow_draft_variable.py |
| PATCH, GET, DELETE | `/console/api/apps/<uuid:app_id>/workflows/draft/variables/<uuid:variable_id>` | `variableapi` | controllers.console.app.workflow_draft_variable.py |
| PUT | `/console/api/apps/<uuid:app_id>/workflows/draft/variables/<uuid:variable_id>/reset` | `variableresetapi` | controllers.console.app.workflow_draft_variable.py |
| GET | `/console/api/apps/<uuid:app_id>/workflows/draft/conversation-variables` | `conversationvariablecollectionapi` | controllers.console.app.workflow_draft_variable.py |
| GET | `/console/api/apps/<uuid:app_id>/workflows/draft/system-variables` | `systemvariablecollectionapi` | controllers.console.app.workflow_draft_variable.py |
| GET | `/console/api/apps/<uuid:app_id>/workflows/draft/environment-variables` | `environmentvariablecollectionapi` | controllers.console.app.workflow_draft_variable.py |
| GET | `/console/api/apps/<uuid:app_id>/advanced-chat/workflow-runs` | `advancedchatappworkflowrunlistapi` | controllers.console.app.workflow_run.py |
| GET | `/console/api/apps/<uuid:app_id>/workflow-runs` | `workflowrunlistapi` | controllers.console.app.workflow_run.py |
| GET | `/console/api/apps/<uuid:app_id>/workflow-runs/<uuid:run_id>` | `workflowrundetailapi` | controllers.console.app.workflow_run.py |
| GET | `/console/api/apps/<uuid:app_id>/workflow-runs/<uuid:run_id>/node-executions` | `workflowrunnodeexecutionlistapi` | controllers.console.app.workflow_run.py |
| GET | `/console/api/apps/<uuid:app_id>/workflow/statistics/daily-conversations` | `workflowdailyrunsstatistic` | controllers.console.app.workflow_statistic.py |
| GET | `/console/api/apps/<uuid:app_id>/workflow/statistics/daily-terminals` | `workflowdailyterminalsstatistic` | controllers.console.app.workflow_statistic.py |
| GET | `/console/api/apps/<uuid:app_id>/workflow/statistics/token-costs` | `workflowdailytokencoststatistic` | controllers.console.app.workflow_statistic.py |
| GET | `/console/api/apps/<uuid:app_id>/workflow/statistics/average-app-interactions` | `workflowaverageappinteractionstatistic` | controllers.console.app.workflow_statistic.py |
| GET | `/console/api/activate/check` | `activatecheckapi` | controllers.console.auth.activate.py |
| POST | `/console/api/activate` | `activateapi` | controllers.console.auth.activate.py |
| GET | `/console/api/api-key-auth/data-source` | `apikeyauthdatasource` | controllers.console.auth.data_source_bearer_auth.py |
| POST | `/console/api/api-key-auth/data-source/binding` | `apikeyauthdatasourcebinding` | controllers.console.auth.data_source_bearer_auth.py |
| DELETE | `/console/api/api-key-auth/data-source/<uuid:binding_id>` | `apikeyauthdatasourcebindingdelete` | controllers.console.auth.data_source_bearer_auth.py |
| GET | `/console/api/oauth/data-source/<string:provider>` | `oauthdatasource` | controllers.console.auth.data_source_oauth.py |
| GET | `/console/api/oauth/data-source/callback/<string:provider>` | `oauthdatasourcecallback` | controllers.console.auth.data_source_oauth.py |
| GET | `/console/api/oauth/data-source/binding/<string:provider>` | `oauthdatasourcebinding` | controllers.console.auth.data_source_oauth.py |
| GET | `/console/api/oauth/data-source/<string:provider>/<uuid:binding_id>/sync` | `oauthdatasourcesync` | controllers.console.auth.data_source_oauth.py |
| POST | `/console/api/forgot-password` | `forgotpasswordsendemailapi` | controllers.console.auth.forgot_password.py |
| POST | `/console/api/forgot-password/validity` | `forgotpasswordcheckapi` | controllers.console.auth.forgot_password.py |
| POST | `/console/api/forgot-password/resets` | `forgotpasswordresetapi` | controllers.console.auth.forgot_password.py |
| POST | `/console/api/login` | `loginapi` | controllers.console.auth.login.py |
| GET | `/console/api/logout` | `logoutapi` | controllers.console.auth.login.py |
| POST | `/console/api/email-code-login` | `emailcodeloginsendemailapi` | controllers.console.auth.login.py |
| POST | `/console/api/email-code-login/validity` | `emailcodeloginapi` | controllers.console.auth.login.py |
| POST | `/console/api/reset-password` | `resetpasswordsendemailapi` | controllers.console.auth.login.py |
| POST | `/console/api/refresh-token` | `refreshtokenapi` | controllers.console.auth.login.py |
| GET | `/console/api/oauth/login/<provider>` | `oauthlogin` | controllers.console.auth.oauth.py |
| GET | `/console/api/oauth/authorize/<provider>` | `oauthcallback` | controllers.console.auth.oauth.py |
| GET | `/console/api/billing/subscription` | `subscription` | controllers.console.billing.billing.py |
| GET | `/console/api/billing/invoices` | `invoices` | controllers.console.billing.billing.py |
| GET | `/console/api/compliance/download` | `complianceapi` | controllers.console.billing.compliance.py |
| PATCH, GET | `/console/api/data-source/integrates/<uuid:binding_id>/<string:action>` | `datasourceapi` | controllers.console.datasets.data_source.py |
| PATCH, GET | `/console/api/data-source/integrates` | `datasourceapi` | controllers.console.datasets.data_source.py |
| GET | `/console/api/notion/pre-import/pages` | `datasourcenotionlistapi` | controllers.console.datasets.data_source.py |
| POST, GET | `/console/api/notion/workspaces/<uuid:workspace_id>/pages/<uuid:page_id>/<string:page_type>/preview` | `datasourcenotionapi` | controllers.console.datasets.data_source.py |
| POST, GET | `/console/api/datasets/notion-indexing-estimate` | `datasourcenotionapi` | controllers.console.datasets.data_source.py |
| GET | `/console/api/datasets/<uuid:dataset_id>/notion/sync` | `datasourcenotiondatasetsyncapi` | controllers.console.datasets.data_source.py |
| GET | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/notion/sync` | `datasourcenotiondocumentsyncapi` | controllers.console.datasets.data_source.py |
| POST, GET | `/console/api/datasets` | `datasetlistapi` | controllers.console.datasets.datasets.py |
| PATCH, GET, DELETE | `/console/api/datasets/<uuid:dataset_id>` | `datasetapi` | controllers.console.datasets.datasets.py |
| GET | `/console/api/datasets/<uuid:dataset_id>/use-check` | `datasetusecheckapi` | controllers.console.datasets.datasets.py |
| GET | `/console/api/datasets/<uuid:dataset_id>/queries` | `datasetqueryapi` | controllers.console.datasets.datasets.py |
| GET | `/console/api/datasets/<uuid:dataset_id>/error-docs` | `dataseterrordocs` | controllers.console.datasets.datasets.py |
| POST | `/console/api/datasets/indexing-estimate` | `datasetindexingestimateapi` | controllers.console.datasets.datasets.py |
| GET | `/console/api/datasets/<uuid:dataset_id>/related-apps` | `datasetrelatedapplistapi` | controllers.console.datasets.datasets.py |
| GET | `/console/api/datasets/<uuid:dataset_id>/indexing-status` | `datasetindexingstatusapi` | controllers.console.datasets.datasets.py |
| POST, GET | `/console/api/datasets/api-keys` | `datasetapikeyapi` | controllers.console.datasets.datasets.py |
| DELETE | `/console/api/datasets/api-keys/<uuid:api_key_id>` | `datasetapideleteapi` | controllers.console.datasets.datasets.py |
| GET | `/console/api/datasets/api-base-info` | `datasetapibaseurlapi` | controllers.console.datasets.datasets.py |
| GET | `/console/api/datasets/retrieval-setting` | `datasetretrievalsettingapi` | controllers.console.datasets.datasets.py |
| GET | `/console/api/datasets/retrieval-setting/<string:vector_type>` | `datasetretrievalsettingmockapi` | controllers.console.datasets.datasets.py |
| GET | `/console/api/datasets/<uuid:dataset_id>/permission-part-users` | `datasetpermissionuserlistapi` | controllers.console.datasets.datasets.py |
| GET | `/console/api/datasets/<uuid:dataset_id>/auto-disable-logs` | `datasetautodisablelogapi` | controllers.console.datasets.datasets.py |
| GET | `/console/api/datasets/process-rule` | `getprocessruleapi` | controllers.console.datasets.datasets_document.py |
| POST, GET, DELETE | `/console/api/datasets/<uuid:dataset_id>/documents` | `datasetdocumentlistapi` | controllers.console.datasets.datasets_document.py |
| POST | `/console/api/datasets/init` | `datasetinitapi` | controllers.console.datasets.datasets_document.py |
| GET | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/indexing-estimate` | `documentindexingestimateapi` | controllers.console.datasets.datasets_document.py |
| GET | `/console/api/datasets/<uuid:dataset_id>/batch/<string:batch>/indexing-estimate` | `documentbatchindexingestimateapi` | controllers.console.datasets.datasets_document.py |
| GET | `/console/api/datasets/<uuid:dataset_id>/batch/<string:batch>/indexing-status` | `documentbatchindexingstatusapi` | controllers.console.datasets.datasets_document.py |
| GET | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/indexing-status` | `documentindexingstatusapi` | controllers.console.datasets.datasets_document.py |
| GET | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>` | `documentdetailapi` | controllers.console.datasets.datasets_document.py |
| PATCH | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/processing/<string:action>` | `documentprocessingapi` | controllers.console.datasets.datasets_document.py |
| DELETE | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>` | `documentdeleteapi` | controllers.console.datasets.datasets_document.py |
| PUT | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/metadata` | `documentmetadataapi` | controllers.console.datasets.datasets_document.py |
| PATCH | `/console/api/datasets/<uuid:dataset_id>/documents/status/<string:action>/batch` | `documentstatusapi` | controllers.console.datasets.datasets_document.py |
| PATCH | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/processing/pause` | `documentpauseapi` | controllers.console.datasets.datasets_document.py |
| PATCH | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/processing/resume` | `documentrecoverapi` | controllers.console.datasets.datasets_document.py |
| POST | `/console/api/datasets/<uuid:dataset_id>/retry` | `documentretryapi` | controllers.console.datasets.datasets_document.py |
| POST | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/rename` | `documentrenameapi` | controllers.console.datasets.datasets_document.py |
| GET | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/website-sync` | `websitedocumentsyncapi` | controllers.console.datasets.datasets_document.py |
| GET, DELETE | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/segments` | `datasetdocumentsegmentlistapi` | controllers.console.datasets.datasets_segments.py |
| PATCH | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/segment/<string:action>` | `datasetdocumentsegmentapi` | controllers.console.datasets.datasets_segments.py |
| POST | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/segment` | `datasetdocumentsegmentaddapi` | controllers.console.datasets.datasets_segments.py |
| PATCH, DELETE | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/segments/<uuid:segment_id>` | `datasetdocumentsegmentupdateapi` | controllers.console.datasets.datasets_segments.py |
| POST, GET | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/segments/batch_import` | `datasetdocumentsegmentbatchimportapi` | controllers.console.datasets.datasets_segments.py |
| POST, GET | `/console/api/datasets/batch_import_status/<uuid:job_id>` | `datasetdocumentsegmentbatchimportapi` | controllers.console.datasets.datasets_segments.py |
| PATCH, POST, GET | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/segments/<uuid:segment_id>/child_chunks` | `childchunkaddapi` | controllers.console.datasets.datasets_segments.py |
| PATCH, DELETE | `/console/api/datasets/<uuid:dataset_id>/documents/<uuid:document_id>/segments/<uuid:segment_id>/child_chunks/<uuid:child_chunk_id>` | `childchunkupdateapi` | controllers.console.datasets.datasets_segments.py |
| POST | `/console/api/datasets/<uuid:dataset_id>/external-hit-testing` | `externalknowledgehittestingapi` | controllers.console.datasets.external.py |
| POST | `/console/api/datasets/external` | `externaldatasetcreateapi` | controllers.console.datasets.external.py |
| POST, GET | `/console/api/datasets/external-knowledge-api` | `externalapitemplatelistapi` | controllers.console.datasets.external.py |
| PATCH, GET, DELETE | `/console/api/datasets/external-knowledge-api/<uuid:external_knowledge_api_id>` | `externalapitemplateapi` | controllers.console.datasets.external.py |
| GET | `/console/api/datasets/external-knowledge-api/<uuid:external_knowledge_api_id>/use-check` | `externalapiusecheckapi` | controllers.console.datasets.external.py |
| POST | `/console/api/test/retrieval` | `bedrockretrievalapi` | controllers.console.datasets.external.py |
| POST | `/console/api/datasets/<uuid:dataset_id>/hit-testing` | `hittestingapi` | controllers.console.datasets.hit_testing.py |
| POST, GET | `/console/api/datasets/<uuid:dataset_id>/metadata` | `datasetmetadatacreateapi` | controllers.console.datasets.metadata.py |
| PATCH, DELETE | `/console/api/datasets/<uuid:dataset_id>/metadata/<uuid:metadata_id>` | `datasetmetadataapi` | controllers.console.datasets.metadata.py |
| GET | `/console/api/datasets/metadata/built-in` | `datasetmetadatabuiltinfieldapi` | controllers.console.datasets.metadata.py |
| POST | `/console/api/datasets/<uuid:dataset_id>/metadata/built-in/<string:action>` | `datasetmetadatabuiltinfieldactionapi` | controllers.console.datasets.metadata.py |
| POST | `/console/api/datasets/<uuid:dataset_id>/documents/metadata` | `documentmetadataeditapi` | controllers.console.datasets.metadata.py |
| POST | `/console/api/website/crawl` | `websitecrawlapi` | controllers.console.datasets.website.py |
| GET | `/console/api/website/crawl/status/<string:job_id>` | `websitecrawlstatusapi` | controllers.console.datasets.website.py |
| POST, GET | `/console/api/installed-apps` | `installedappslistapi` | controllers.console.explore.installed_app.py |
| PATCH, DELETE | `/console/api/installed-apps/<uuid:installed_app_id>` | `installedappapi` | controllers.console.explore.installed_app.py |
| GET | `/console/api/installed-apps/<uuid:installed_app_id>/parameters` | `installed_app_parameters` | controllers.console.explore.parameter.py |
| GET | `/console/api/installed-apps/<uuid:installed_app_id>/meta` | `installed_app_meta` | controllers.console.explore.parameter.py |
| GET | `/console/api/explore/apps` | `recommendedapplistapi` | controllers.console.explore.recommended_app.py |
| GET | `/console/api/explore/apps/<uuid:app_id>` | `recommendedappapi` | controllers.console.explore.recommended_app.py |
| POST, GET | `/console/api/installed-apps/<uuid:installed_app_id>/saved-messages` | `installed_app_saved_messages` | controllers.console.explore.saved_message.py |
| DELETE | `/console/api/installed-apps/<uuid:installed_app_id>/saved-messages/<uuid:message_id>` | `installed_app_saved_message` | controllers.console.explore.saved_message.py |
| POST | `/console/api/installed-apps/<uuid:installed_app_id>/audio-to-text` | `installed_app_audio` | controllers.console.explore.audio.py |
| POST | `/console/api/installed-apps/<uuid:installed_app_id>/text-to-audio` | `installed_app_text` | controllers.console.explore.audio.py |
| POST | `/console/api/installed-apps/<uuid:installed_app_id>/completion-messages` | `installed_app_completion` | controllers.console.explore.completion.py |
| POST | `/console/api/installed-apps/<uuid:installed_app_id>/completion-messages/<string:task_id>/stop` | `installed_app_stop_completion` | controllers.console.explore.completion.py |
| POST | `/console/api/installed-apps/<uuid:installed_app_id>/chat-messages` | `installed_app_chat_completion` | controllers.console.explore.completion.py |
| POST | `/console/api/installed-apps/<uuid:installed_app_id>/chat-messages/<string:task_id>/stop` | `installed_app_stop_chat_completion` | controllers.console.explore.completion.py |
| POST | `/console/api/installed-apps/<uuid:installed_app_id>/conversations/<uuid:c_id>/name` | `installed_app_conversation_rename` | controllers.console.explore.conversation.py |
| GET | `/console/api/installed-apps/<uuid:installed_app_id>/conversations` | `installed_app_conversations` | controllers.console.explore.conversation.py |
| DELETE | `/console/api/installed-apps/<uuid:installed_app_id>/conversations/<uuid:c_id>` | `installed_app_conversation` | controllers.console.explore.conversation.py |
| PATCH | `/console/api/installed-apps/<uuid:installed_app_id>/conversations/<uuid:c_id>/pin` | `installed_app_conversation_pin` | controllers.console.explore.conversation.py |
| PATCH | `/console/api/installed-apps/<uuid:installed_app_id>/conversations/<uuid:c_id>/unpin` | `installed_app_conversation_unpin` | controllers.console.explore.conversation.py |
| GET | `/console/api/installed-apps/<uuid:installed_app_id>/messages` | `installed_app_messages` | controllers.console.explore.message.py |
| POST | `/console/api/installed-apps/<uuid:installed_app_id>/messages/<uuid:message_id>/feedbacks` | `installed_app_message_feedback` | controllers.console.explore.message.py |
| GET | `/console/api/installed-apps/<uuid:installed_app_id>/messages/<uuid:message_id>/more-like-this` | `installed_app_more_like_this` | controllers.console.explore.message.py |
| GET | `/console/api/installed-apps/<uuid:installed_app_id>/messages/<uuid:message_id>/suggested-questions` | `installed_app_suggested_question` | controllers.console.explore.message.py |
| POST | `/console/api/installed-apps/<uuid:installed_app_id>/workflows/run` | `installedappworkflowrunapi` | controllers.console.explore.workflow.py |
| POST | `/console/api/installed-apps/<uuid:installed_app_id>/workflows/tasks/<string:task_id>/stop` | `installedappworkflowtaskstopapi` | controllers.console.explore.workflow.py |
| POST, GET | `/console/api/tags` | `taglistapi` | controllers.console.tag.tags.py |
| PATCH, DELETE | `/console/api/tags/<uuid:tag_id>` | `tagupdatedeleteapi` | controllers.console.tag.tags.py |
| POST | `/console/api/tag-bindings/create` | `tagbindingcreateapi` | controllers.console.tag.tags.py |
| POST | `/console/api/tag-bindings/remove` | `tagbindingdeleteapi` | controllers.console.tag.tags.py |
| POST | `/console/api/account/init` | `accountinitapi` | controllers.console.workspace.account.py |
| GET | `/console/api/account/profile` | `accountprofileapi` | controllers.console.workspace.account.py |
| POST | `/console/api/account/name` | `accountnameapi` | controllers.console.workspace.account.py |
| POST | `/console/api/account/avatar` | `accountavatarapi` | controllers.console.workspace.account.py |
| POST | `/console/api/account/interface-language` | `accountinterfacelanguageapi` | controllers.console.workspace.account.py |
| POST | `/console/api/account/interface-theme` | `accountinterfacethemeapi` | controllers.console.workspace.account.py |
| POST | `/console/api/account/timezone` | `accounttimezoneapi` | controllers.console.workspace.account.py |
| POST | `/console/api/account/password` | `accountpasswordapi` | controllers.console.workspace.account.py |
| GET | `/console/api/account/integrates` | `accountintegrateapi` | controllers.console.workspace.account.py |
| GET | `/console/api/account/delete/verify` | `accountdeleteverifyapi` | controllers.console.workspace.account.py |
| POST | `/console/api/account/delete` | `accountdeleteapi` | controllers.console.workspace.account.py |
| POST | `/console/api/account/delete/feedback` | `accountdeleteupdatefeedbackapi` | controllers.console.workspace.account.py |
| GET | `/console/api/account/education/verify` | `educationverifyapi` | controllers.console.workspace.account.py |
| POST, GET | `/console/api/account/education` | `educationapi` | controllers.console.workspace.account.py |
| GET | `/console/api/account/education/autocomplete` | `educationautocompleteapi` | controllers.console.workspace.account.py |
| POST | `/console/api/account/change-email` | `changeemailsendemailapi` | controllers.console.workspace.account.py |
| POST | `/console/api/account/change-email/validity` | `changeemailcheckapi` | controllers.console.workspace.account.py |
| POST | `/console/api/account/change-email/reset` | `changeemailresetapi` | controllers.console.workspace.account.py |
| POST | `/console/api/account/change-email/check-email-unique` | `checkemailunique` | controllers.console.workspace.account.py |
| GET | `/console/api/workspaces/current/agent-providers` | `agentproviderlistapi` | controllers.console.workspace.agent_providers.py |
| GET | `/console/api/workspaces/current/agent-provider/<path:provider_name>` | `agentproviderapi` | controllers.console.workspace.agent_providers.py |
| POST | `/console/api/workspaces/current/endpoints/create` | `endpointcreateapi` | controllers.console.workspace.endpoint.py |
| GET | `/console/api/workspaces/current/endpoints/list` | `endpointlistapi` | controllers.console.workspace.endpoint.py |
| GET | `/console/api/workspaces/current/endpoints/list/plugin` | `endpointlistforsinglepluginapi` | controllers.console.workspace.endpoint.py |
| POST | `/console/api/workspaces/current/endpoints/delete` | `endpointdeleteapi` | controllers.console.workspace.endpoint.py |
| POST | `/console/api/workspaces/current/endpoints/update` | `endpointupdateapi` | controllers.console.workspace.endpoint.py |
| POST | `/console/api/workspaces/current/endpoints/enable` | `endpointenableapi` | controllers.console.workspace.endpoint.py |
| POST | `/console/api/workspaces/current/endpoints/disable` | `endpointdisableapi` | controllers.console.workspace.endpoint.py |
| POST | `/console/api/workspaces/current/model-providers/<path:provider>/models/load-balancing-configs/credentials-validate` | `loadbalancingcredentialsvalidateapi` | controllers.console.workspace.load_balancing_config.py |
| POST | `/console/api/workspaces/current/model-providers/<path:provider>/models/load-balancing-configs/<string:config_id>/credentials-validate` | `loadbalancingconfigcredentialsvalidateapi` | controllers.console.workspace.load_balancing_config.py |
| GET | `/console/api/workspaces/current/members` | `memberlistapi` | controllers.console.workspace.members.py |
| POST | `/console/api/workspaces/current/members/invite-email` | `memberinviteemailapi` | controllers.console.workspace.members.py |
| DELETE | `/console/api/workspaces/current/members/<uuid:member_id>` | `membercancelinviteapi` | controllers.console.workspace.members.py |
| PUT | `/console/api/workspaces/current/members/<uuid:member_id>/update-role` | `memberupdateroleapi` | controllers.console.workspace.members.py |
| GET | `/console/api/workspaces/current/dataset-operators` | `datasetoperatormemberlistapi` | controllers.console.workspace.members.py |
| POST | `/console/api/workspaces/current/members/send-owner-transfer-confirm-email` | `sendownertransferemailapi` | controllers.console.workspace.members.py |
| POST | `/console/api/workspaces/current/members/owner-transfer-check` | `ownertransfercheckapi` | controllers.console.workspace.members.py |
| POST | `/console/api/workspaces/current/members/<uuid:member_id>/owner-transfer` | `ownertransfer` | controllers.console.workspace.members.py |
| GET | `/console/api/workspaces/current/model-providers` | `modelproviderlistapi` | controllers.console.workspace.model_providers.py |
| GET | `/console/api/workspaces/current/model-providers/<path:provider>/credentials` | `modelprovidercredentialapi` | controllers.console.workspace.model_providers.py |
| POST | `/console/api/workspaces/current/model-providers/<path:provider>/credentials/validate` | `modelprovidervalidateapi` | controllers.console.workspace.model_providers.py |
| POST, DELETE | `/console/api/workspaces/current/model-providers/<path:provider>` | `modelproviderapi` | controllers.console.workspace.model_providers.py |
| POST | `/console/api/workspaces/current/model-providers/<path:provider>/preferred-provider-type` | `preferredprovidertypeupdateapi` | controllers.console.workspace.model_providers.py |
| GET | `/console/api/workspaces/current/model-providers/<path:provider>/checkout-url` | `modelproviderpaymentcheckouturlapi` | controllers.console.workspace.model_providers.py |
| GET | `/console/api/workspaces/<string:tenant_id>/model-providers/<path:provider>/<string:icon_type>/<string:lang>` | `modelprovidericonapi` | controllers.console.workspace.model_providers.py |
| POST, GET, DELETE | `/console/api/workspaces/current/model-providers/<path:provider>/models` | `modelprovidermodelapi` | controllers.console.workspace.models.py |
| PATCH | `/console/api/workspaces/current/model-providers/<path:provider>/models/enable` | `model-provider-model-enable` | controllers.console.workspace.models.py |
| PATCH | `/console/api/workspaces/current/model-providers/<path:provider>/models/disable` | `model-provider-model-disable` | controllers.console.workspace.models.py |
| GET | `/console/api/workspaces/current/model-providers/<path:provider>/models/credentials` | `modelprovidermodelcredentialapi` | controllers.console.workspace.models.py |
| POST | `/console/api/workspaces/current/model-providers/<path:provider>/models/credentials/validate` | `modelprovidermodelvalidateapi` | controllers.console.workspace.models.py |
| GET | `/console/api/workspaces/current/model-providers/<path:provider>/models/parameter-rules` | `modelprovidermodelparameterruleapi` | controllers.console.workspace.models.py |
| GET | `/console/api/workspaces/current/models/model-types/<string:model_type>` | `modelprovideravailablemodelapi` | controllers.console.workspace.models.py |
| POST, GET | `/console/api/workspaces/current/default-model` | `defaultmodelapi` | controllers.console.workspace.models.py |
| GET | `/console/api/workspaces/current/plugin/debugging-key` | `plugindebuggingkeyapi` | controllers.console.workspace.plugin.py |
| GET | `/console/api/workspaces/current/plugin/list` | `pluginlistapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/list/latest-versions` | `pluginlistlatestversionsapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/list/installations/ids` | `pluginlistinstallationsfromidsapi` | controllers.console.workspace.plugin.py |
| GET | `/console/api/workspaces/current/plugin/icon` | `pluginiconapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/upload/pkg` | `pluginuploadfrompkgapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/upload/github` | `pluginuploadfromgithubapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/upload/bundle` | `pluginuploadfrombundleapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/install/pkg` | `plugininstallfrompkgapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/install/github` | `plugininstallfromgithubapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/upgrade/marketplace` | `pluginupgradefrommarketplaceapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/upgrade/github` | `pluginupgradefromgithubapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/install/marketplace` | `plugininstallfrommarketplaceapi` | controllers.console.workspace.plugin.py |
| GET | `/console/api/workspaces/current/plugin/fetch-manifest` | `pluginfetchmanifestapi` | controllers.console.workspace.plugin.py |
| GET | `/console/api/workspaces/current/plugin/tasks` | `pluginfetchinstalltasksapi` | controllers.console.workspace.plugin.py |
| GET | `/console/api/workspaces/current/plugin/tasks/<task_id>` | `pluginfetchinstalltaskapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/tasks/<task_id>/delete` | `plugindeleteinstalltaskapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/tasks/delete_all` | `plugindeleteallinstalltaskitemsapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/tasks/<task_id>/delete/<path:identifier>` | `plugindeleteinstalltaskitemapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/uninstall` | `pluginuninstallapi` | controllers.console.workspace.plugin.py |
| GET | `/console/api/workspaces/current/plugin/marketplace/pkg` | `pluginfetchmarketplacepkgapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/permission/change` | `pluginchangepermissionapi` | controllers.console.workspace.plugin.py |
| GET | `/console/api/workspaces/current/plugin/permission/fetch` | `pluginfetchpermissionapi` | controllers.console.workspace.plugin.py |
| GET | `/console/api/workspaces/current/plugin/parameters/dynamic-options` | `pluginfetchdynamicselectoptionsapi` | controllers.console.workspace.plugin.py |
| GET | `/console/api/workspaces/current/plugin/preferences/fetch` | `pluginfetchpreferencesapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/preferences/change` | `pluginchangepreferencesapi` | controllers.console.workspace.plugin.py |
| POST | `/console/api/workspaces/current/plugin/preferences/autoupgrade/exclude` | `pluginautoupgradeexcludepluginapi` | controllers.console.workspace.plugin.py |
| GET | `/console/api/workspaces/current/tool-providers` | `toolproviderlistapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/oauth/plugin/<path:provider>/tool/authorization-url` | `toolpluginoauthapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/oauth/plugin/<path:provider>/tool/callback` | `tooloauthcallback` | controllers.console.workspace.tool_providers.py |
| POST, GET, DELETE | `/console/api/workspaces/current/tool-provider/builtin/<path:provider>/oauth/custom-client` | `tooloauthcustomclient` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tool-provider/builtin/<path:provider>/tools` | `toolbuiltinproviderlisttoolsapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tool-provider/builtin/<path:provider>/info` | `toolbuiltinproviderinfoapi` | controllers.console.workspace.tool_providers.py |
| POST | `/console/api/workspaces/current/tool-provider/builtin/<path:provider>/add` | `toolbuiltinprovideraddapi` | controllers.console.workspace.tool_providers.py |
| POST | `/console/api/workspaces/current/tool-provider/builtin/<path:provider>/delete` | `toolbuiltinproviderdeleteapi` | controllers.console.workspace.tool_providers.py |
| POST | `/console/api/workspaces/current/tool-provider/builtin/<path:provider>/update` | `toolbuiltinproviderupdateapi` | controllers.console.workspace.tool_providers.py |
| POST | `/console/api/workspaces/current/tool-provider/builtin/<path:provider>/default-credential` | `toolbuiltinprovidersetdefaultapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tool-provider/builtin/<path:provider>/credential/info` | `toolbuiltinprovidergetcredentialinfoapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tool-provider/builtin/<path:provider>/credentials` | `toolbuiltinprovidergetcredentialsapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tool-provider/builtin/<path:provider>/credential/schema/<path:credential_type>` | `toolbuiltinprovidercredentialsschemaapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tool-provider/builtin/<path:provider>/oauth/client-schema` | `toolbuiltinprovidergetoauthclientschemaapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tool-provider/builtin/<path:provider>/icon` | `toolbuiltinprovidericonapi` | controllers.console.workspace.tool_providers.py |
| POST | `/console/api/workspaces/current/tool-provider/api/add` | `toolapiprovideraddapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tool-provider/api/remote` | `toolapiprovidergetremoteschemaapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tool-provider/api/tools` | `toolapiproviderlisttoolsapi` | controllers.console.workspace.tool_providers.py |
| POST | `/console/api/workspaces/current/tool-provider/api/update` | `toolapiproviderupdateapi` | controllers.console.workspace.tool_providers.py |
| POST | `/console/api/workspaces/current/tool-provider/api/delete` | `toolapiproviderdeleteapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tool-provider/api/get` | `toolapiprovidergetapi` | controllers.console.workspace.tool_providers.py |
| POST | `/console/api/workspaces/current/tool-provider/api/schema` | `toolapiproviderschemaapi` | controllers.console.workspace.tool_providers.py |
| POST | `/console/api/workspaces/current/tool-provider/api/test/pre` | `toolapiproviderprevioustestapi` | controllers.console.workspace.tool_providers.py |
| POST | `/console/api/workspaces/current/tool-provider/workflow/create` | `toolworkflowprovidercreateapi` | controllers.console.workspace.tool_providers.py |
| POST | `/console/api/workspaces/current/tool-provider/workflow/update` | `toolworkflowproviderupdateapi` | controllers.console.workspace.tool_providers.py |
| POST | `/console/api/workspaces/current/tool-provider/workflow/delete` | `toolworkflowproviderdeleteapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tool-provider/workflow/get` | `toolworkflowprovidergetapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tool-provider/workflow/tools` | `toolworkflowproviderlisttoolapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tool-provider/mcp/tools/<path:provider_id>` | `toolmcpdetailapi` | controllers.console.workspace.tool_providers.py |
| POST, PUT, DELETE | `/console/api/workspaces/current/tool-provider/mcp` | `toolprovidermcpapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tool-provider/mcp/update/<path:provider_id>` | `toolmcpupdateapi` | controllers.console.workspace.tool_providers.py |
| POST | `/console/api/workspaces/current/tool-provider/mcp/auth` | `toolmcpauthapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/mcp/oauth/callback` | `toolmcpcallbackapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tools/builtin` | `toolbuiltinlistapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tools/api` | `toolapilistapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tools/mcp` | `toolmcplistallapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tools/workflow` | `toolworkflowlistapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces/current/tool-labels` | `toollabelsapi` | controllers.console.workspace.tool_providers.py |
| GET | `/console/api/workspaces` | `tenantlistapi` | controllers.console.workspace.workspace.py |
| GET | `/console/api/all-workspaces` | `workspacelistapi` | controllers.console.workspace.workspace.py |
| GET | `/console/api/workspaces/current` | `workspaces_current` | controllers.console.workspace.workspace.py |
| GET | `/console/api/info` | `info` | controllers.console.workspace.workspace.py |
| POST | `/console/api/workspaces/switch` | `switchworkspaceapi` | controllers.console.workspace.workspace.py |
| POST | `/console/api/workspaces/custom-config` | `customconfigworkspaceapi` | controllers.console.workspace.workspace.py |
| POST | `/console/api/workspaces/custom-config/webapp-logo/upload` | `webapplogoworkspaceapi` | controllers.console.workspace.workspace.py |
| POST | `/console/api/workspaces/info` | `workspaceinfoapi` | controllers.console.workspace.workspace.py |

## files

| Method | URL | Function | file |
|--------|-----|----------|-------------|
| GET | `/files/<uuid:file_id>/image-preview` | `imagepreviewapi` | controllers.files.image_preview.py |
| GET | `/files/<uuid:file_id>/file-preview` | `filepreviewapi` | controllers.files.image_preview.py |
| GET | `/files/workspaces/<uuid:workspace_id>/webapp-logo` | `workspacewebapplogoapi` | controllers.files.image_preview.py |
| GET | `/files/tools/<uuid:file_id>.<string:extension>` | `toolfilepreviewapi` | controllers.files.tool_files.py |
| POST | `/files/upload/for-plugin` | `pluginuploadfileapi` | controllers.files.upload.py |

## inner_api
**URL Prefix:** `/inner/api`

| Method | URL | Function | file |
|--------|-----|----------|-------------|
| POST | `/inner/api/enterprise/mail` | `enterprisemail` | controllers.inner_api.mail.py |
| POST | `/inner/api/invoke/llm` | `plugininvokellmapi` | controllers.inner_api.plugin.plugin.py |
| POST | `/inner/api/invoke/llm/structured-output` | `plugininvokellmwithstructuredoutputapi` | controllers.inner_api.plugin.plugin.py |
| POST | `/inner/api/invoke/text-embedding` | `plugininvoketextembeddingapi` | controllers.inner_api.plugin.plugin.py |
| POST | `/inner/api/invoke/rerank` | `plugininvokererankapi` | controllers.inner_api.plugin.plugin.py |
| POST | `/inner/api/invoke/tts` | `plugininvokettsapi` | controllers.inner_api.plugin.plugin.py |
| POST | `/inner/api/invoke/speech2text` | `plugininvokespeech2textapi` | controllers.inner_api.plugin.plugin.py |
| POST | `/inner/api/invoke/moderation` | `plugininvokemoderationapi` | controllers.inner_api.plugin.plugin.py |
| POST | `/inner/api/invoke/tool` | `plugininvoketoolapi` | controllers.inner_api.plugin.plugin.py |
| POST | `/inner/api/invoke/parameter-extractor` | `plugininvokeparameterextractornodeapi` | controllers.inner_api.plugin.plugin.py |
| POST | `/inner/api/invoke/question-classifier` | `plugininvokequestionclassifiernodeapi` | controllers.inner_api.plugin.plugin.py |
| POST | `/inner/api/invoke/app` | `plugininvokeappapi` | controllers.inner_api.plugin.plugin.py |
| POST | `/inner/api/invoke/encrypt` | `plugininvokeencryptapi` | controllers.inner_api.plugin.plugin.py |
| POST | `/inner/api/invoke/summary` | `plugininvokesummaryapi` | controllers.inner_api.plugin.plugin.py |
| POST | `/inner/api/upload/file/request` | `pluginuploadfilerequestapi` | controllers.inner_api.plugin.plugin.py |
| POST | `/inner/api/fetch/app/info` | `pluginfetchappinfoapi` | controllers.inner_api.plugin.plugin.py |
| POST | `/inner/api/enterprise/workspace` | `enterpriseworkspace` | controllers.inner_api.workspace.workspace.py |
| POST | `/inner/api/enterprise/workspace/ownerless` | `enterpriseworkspacenoowneremail` | controllers.inner_api.workspace.workspace.py |

## mcp
**URL Prefix:** `/mcp`

| Method | URL | Function | file |
|--------|-----|----------|-------------|
| POST | `/mcp/server/<string:server_code>/mcp` | `mcpappapi` | controllers.mcp.mcp.py |
