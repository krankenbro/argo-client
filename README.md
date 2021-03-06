# Go API client for swagger

Description of all APIs

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: version not set
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountServiceApi* | [**AccountServiceCanI**](docs/AccountServiceApi.md#accountservicecani) | **Get** /api/v1/account/can-i/{resource}/{action}/{subresource} | CanI checks if the current account has permission to perform an action
*AccountServiceApi* | [**AccountServiceCreateToken**](docs/AccountServiceApi.md#accountservicecreatetoken) | **Post** /api/v1/account/{name}/token | CreateToken creates a token
*AccountServiceApi* | [**AccountServiceDeleteToken**](docs/AccountServiceApi.md#accountservicedeletetoken) | **Delete** /api/v1/account/{name}/token/{id} | DeleteToken deletes a token
*AccountServiceApi* | [**AccountServiceGetAccount**](docs/AccountServiceApi.md#accountservicegetaccount) | **Get** /api/v1/account/{name} | GetAccount returns an account
*AccountServiceApi* | [**AccountServiceListAccounts**](docs/AccountServiceApi.md#accountservicelistaccounts) | **Get** /api/v1/account | ListAccounts returns the list of accounts
*AccountServiceApi* | [**AccountServiceUpdatePassword**](docs/AccountServiceApi.md#accountserviceupdatepassword) | **Put** /api/v1/account/password | UpdatePassword updates an account&#39;s password to a new value
*ApplicationServiceApi* | [**ApplicationServiceCreate**](docs/ApplicationServiceApi.md#applicationservicecreate) | **Post** /api/v1/applications | Create creates an application
*ApplicationServiceApi* | [**ApplicationServiceDelete**](docs/ApplicationServiceApi.md#applicationservicedelete) | **Delete** /api/v1/applications/{name} | Delete deletes an application
*ApplicationServiceApi* | [**ApplicationServiceDeleteResource**](docs/ApplicationServiceApi.md#applicationservicedeleteresource) | **Delete** /api/v1/applications/{name}/resource | DeleteResource deletes a single application resource
*ApplicationServiceApi* | [**ApplicationServiceGet**](docs/ApplicationServiceApi.md#applicationserviceget) | **Get** /api/v1/applications/{name} | Get returns an application by name
*ApplicationServiceApi* | [**ApplicationServiceGetApplicationSyncWindows**](docs/ApplicationServiceApi.md#applicationservicegetapplicationsyncwindows) | **Get** /api/v1/applications/{name}/syncwindows | Get returns sync windows of the application
*ApplicationServiceApi* | [**ApplicationServiceGetManifests**](docs/ApplicationServiceApi.md#applicationservicegetmanifests) | **Get** /api/v1/applications/{name}/manifests | GetManifests returns application manifests
*ApplicationServiceApi* | [**ApplicationServiceGetResource**](docs/ApplicationServiceApi.md#applicationservicegetresource) | **Get** /api/v1/applications/{name}/resource | GetResource returns single application resource
*ApplicationServiceApi* | [**ApplicationServiceList**](docs/ApplicationServiceApi.md#applicationservicelist) | **Get** /api/v1/applications | List returns list of applications
*ApplicationServiceApi* | [**ApplicationServiceListResourceActions**](docs/ApplicationServiceApi.md#applicationservicelistresourceactions) | **Get** /api/v1/applications/{name}/resource/actions | ListResourceActions returns list of resource actions
*ApplicationServiceApi* | [**ApplicationServiceListResourceEvents**](docs/ApplicationServiceApi.md#applicationservicelistresourceevents) | **Get** /api/v1/applications/{name}/events | ListResourceEvents returns a list of event resources
*ApplicationServiceApi* | [**ApplicationServiceManagedResources**](docs/ApplicationServiceApi.md#applicationservicemanagedresources) | **Get** /api/v1/applications/{applicationName}/managed-resources | ManagedResources returns list of managed resources
*ApplicationServiceApi* | [**ApplicationServicePatch**](docs/ApplicationServiceApi.md#applicationservicepatch) | **Patch** /api/v1/applications/{name} | Patch patch an application
*ApplicationServiceApi* | [**ApplicationServicePatchResource**](docs/ApplicationServiceApi.md#applicationservicepatchresource) | **Post** /api/v1/applications/{name}/resource | PatchResource patch single application resource
*ApplicationServiceApi* | [**ApplicationServicePodLogs**](docs/ApplicationServiceApi.md#applicationservicepodlogs) | **Get** /api/v1/applications/{name}/pods/{podName}/logs | PodLogs returns stream of log entries for the specified pod. Pod
*ApplicationServiceApi* | [**ApplicationServicePodLogs2**](docs/ApplicationServiceApi.md#applicationservicepodlogs2) | **Get** /api/v1/applications/{name}/logs | PodLogs returns stream of log entries for the specified pod. Pod
*ApplicationServiceApi* | [**ApplicationServiceResourceTree**](docs/ApplicationServiceApi.md#applicationserviceresourcetree) | **Get** /api/v1/applications/{applicationName}/resource-tree | ResourceTree returns resource tree
*ApplicationServiceApi* | [**ApplicationServiceRevisionMetadata**](docs/ApplicationServiceApi.md#applicationservicerevisionmetadata) | **Get** /api/v1/applications/{name}/revisions/{revision}/metadata | Get the meta-data (author, date, tags, message) for a specific revision of the application
*ApplicationServiceApi* | [**ApplicationServiceRollback**](docs/ApplicationServiceApi.md#applicationservicerollback) | **Post** /api/v1/applications/{name}/rollback | Rollback syncs an application to its target state
*ApplicationServiceApi* | [**ApplicationServiceRunResourceAction**](docs/ApplicationServiceApi.md#applicationservicerunresourceaction) | **Post** /api/v1/applications/{name}/resource/actions | RunResourceAction run resource action
*ApplicationServiceApi* | [**ApplicationServiceSync**](docs/ApplicationServiceApi.md#applicationservicesync) | **Post** /api/v1/applications/{name}/sync | Sync syncs an application to its target state
*ApplicationServiceApi* | [**ApplicationServiceTerminateOperation**](docs/ApplicationServiceApi.md#applicationserviceterminateoperation) | **Delete** /api/v1/applications/{name}/operation | TerminateOperation terminates the currently running operation
*ApplicationServiceApi* | [**ApplicationServiceUpdate**](docs/ApplicationServiceApi.md#applicationserviceupdate) | **Put** /api/v1/applications/{application.metadata.name} | Update updates an application
*ApplicationServiceApi* | [**ApplicationServiceUpdateSpec**](docs/ApplicationServiceApi.md#applicationserviceupdatespec) | **Put** /api/v1/applications/{name}/spec | UpdateSpec updates an application spec
*ApplicationServiceApi* | [**ApplicationServiceWatch**](docs/ApplicationServiceApi.md#applicationservicewatch) | **Get** /api/v1/stream/applications | Watch returns stream of application change events
*ApplicationServiceApi* | [**ApplicationServiceWatchResourceTree**](docs/ApplicationServiceApi.md#applicationservicewatchresourcetree) | **Get** /api/v1/stream/applications/{applicationName}/resource-tree | Watch returns stream of application resource tree
*CertificateServiceApi* | [**CertificateServiceCreateCertificate**](docs/CertificateServiceApi.md#certificateservicecreatecertificate) | **Post** /api/v1/certificates | Creates repository certificates on the server
*CertificateServiceApi* | [**CertificateServiceDeleteCertificate**](docs/CertificateServiceApi.md#certificateservicedeletecertificate) | **Delete** /api/v1/certificates | Delete the certificates that match the RepositoryCertificateQuery
*CertificateServiceApi* | [**CertificateServiceListCertificates**](docs/CertificateServiceApi.md#certificateservicelistcertificates) | **Get** /api/v1/certificates | List all available repository certificates
*ClusterServiceApi* | [**ClusterServiceCreate**](docs/ClusterServiceApi.md#clusterservicecreate) | **Post** /api/v1/clusters | Create creates a cluster
*ClusterServiceApi* | [**ClusterServiceDelete**](docs/ClusterServiceApi.md#clusterservicedelete) | **Delete** /api/v1/clusters/{server} | Delete deletes a cluster
*ClusterServiceApi* | [**ClusterServiceGet**](docs/ClusterServiceApi.md#clusterserviceget) | **Get** /api/v1/clusters/{server} | Get returns a cluster by server address
*ClusterServiceApi* | [**ClusterServiceInvalidateCache**](docs/ClusterServiceApi.md#clusterserviceinvalidatecache) | **Post** /api/v1/clusters/{server}/invalidate-cache | InvalidateCache invalidates cluster cache
*ClusterServiceApi* | [**ClusterServiceList**](docs/ClusterServiceApi.md#clusterservicelist) | **Get** /api/v1/clusters | List returns list of clusters
*ClusterServiceApi* | [**ClusterServiceRotateAuth**](docs/ClusterServiceApi.md#clusterservicerotateauth) | **Post** /api/v1/clusters/{server}/rotate-auth | RotateAuth rotates the bearer token used for a cluster
*ClusterServiceApi* | [**ClusterServiceUpdate**](docs/ClusterServiceApi.md#clusterserviceupdate) | **Put** /api/v1/clusters/{cluster.server} | Update updates a cluster
*GPGKeyServiceApi* | [**GPGKeyServiceCreate**](docs/GPGKeyServiceApi.md#gpgkeyservicecreate) | **Post** /api/v1/gpgkeys | Create one or more GPG public keys in the server&#39;s configuration
*GPGKeyServiceApi* | [**GPGKeyServiceDelete**](docs/GPGKeyServiceApi.md#gpgkeyservicedelete) | **Delete** /api/v1/gpgkeys | Delete specified GPG public key from the server&#39;s configuration
*GPGKeyServiceApi* | [**GPGKeyServiceGet**](docs/GPGKeyServiceApi.md#gpgkeyserviceget) | **Get** /api/v1/gpgkeys/{keyID} | Get information about specified GPG public key from the server
*GPGKeyServiceApi* | [**GPGKeyServiceList**](docs/GPGKeyServiceApi.md#gpgkeyservicelist) | **Get** /api/v1/gpgkeys | List all available repository certificates
*ProjectServiceApi* | [**ProjectServiceCreate**](docs/ProjectServiceApi.md#projectservicecreate) | **Post** /api/v1/projects | Create a new project
*ProjectServiceApi* | [**ProjectServiceCreateToken**](docs/ProjectServiceApi.md#projectservicecreatetoken) | **Post** /api/v1/projects/{project}/roles/{role}/token | Create a new project token
*ProjectServiceApi* | [**ProjectServiceDelete**](docs/ProjectServiceApi.md#projectservicedelete) | **Delete** /api/v1/projects/{name} | Delete deletes a project
*ProjectServiceApi* | [**ProjectServiceDeleteToken**](docs/ProjectServiceApi.md#projectservicedeletetoken) | **Delete** /api/v1/projects/{project}/roles/{role}/token/{iat} | Delete a new project token
*ProjectServiceApi* | [**ProjectServiceGet**](docs/ProjectServiceApi.md#projectserviceget) | **Get** /api/v1/projects/{name} | Get returns a project by name
*ProjectServiceApi* | [**ProjectServiceGetDetailedProject**](docs/ProjectServiceApi.md#projectservicegetdetailedproject) | **Get** /api/v1/projects/{name}/detailed | GetDetailedProject returns a project that include project, global project and scoped resources by name
*ProjectServiceApi* | [**ProjectServiceGetGlobalProjects**](docs/ProjectServiceApi.md#projectservicegetglobalprojects) | **Get** /api/v1/projects/{name}/globalprojects | Get returns a virtual project by name
*ProjectServiceApi* | [**ProjectServiceGetSyncWindowsState**](docs/ProjectServiceApi.md#projectservicegetsyncwindowsstate) | **Get** /api/v1/projects/{name}/syncwindows | GetSchedulesState returns true if there are any active sync syncWindows
*ProjectServiceApi* | [**ProjectServiceList**](docs/ProjectServiceApi.md#projectservicelist) | **Get** /api/v1/projects | List returns list of projects
*ProjectServiceApi* | [**ProjectServiceListEvents**](docs/ProjectServiceApi.md#projectservicelistevents) | **Get** /api/v1/projects/{name}/events | ListEvents returns a list of project events
*ProjectServiceApi* | [**ProjectServiceUpdate**](docs/ProjectServiceApi.md#projectserviceupdate) | **Put** /api/v1/projects/{project.metadata.name} | Update updates a project
*RepoCredsServiceApi* | [**RepoCredsServiceCreateRepositoryCredentials**](docs/RepoCredsServiceApi.md#repocredsservicecreaterepositorycredentials) | **Post** /api/v1/repocreds | CreateRepositoryCredentials creates a new repository credential set
*RepoCredsServiceApi* | [**RepoCredsServiceDeleteRepositoryCredentials**](docs/RepoCredsServiceApi.md#repocredsservicedeleterepositorycredentials) | **Delete** /api/v1/repocreds/{url} | DeleteRepositoryCredentials deletes a repository credential set from the configuration
*RepoCredsServiceApi* | [**RepoCredsServiceListRepositoryCredentials**](docs/RepoCredsServiceApi.md#repocredsservicelistrepositorycredentials) | **Get** /api/v1/repocreds | ListRepositoryCredentials gets a list of all configured repository credential sets
*RepoCredsServiceApi* | [**RepoCredsServiceUpdateRepositoryCredentials**](docs/RepoCredsServiceApi.md#repocredsserviceupdaterepositorycredentials) | **Put** /api/v1/repocreds/{creds.url} | UpdateRepositoryCredentials updates a repository credential set
*RepositoryServiceApi* | [**RepositoryServiceCreateRepository**](docs/RepositoryServiceApi.md#repositoryservicecreaterepository) | **Post** /api/v1/repositories | CreateRepository creates a new repository configuration
*RepositoryServiceApi* | [**RepositoryServiceDeleteRepository**](docs/RepositoryServiceApi.md#repositoryservicedeleterepository) | **Delete** /api/v1/repositories/{repo} | DeleteRepository deletes a repository from the configuration
*RepositoryServiceApi* | [**RepositoryServiceGet**](docs/RepositoryServiceApi.md#repositoryserviceget) | **Get** /api/v1/repositories/{repo} | Get returns a repository or its credentials
*RepositoryServiceApi* | [**RepositoryServiceGetAppDetails**](docs/RepositoryServiceApi.md#repositoryservicegetappdetails) | **Post** /api/v1/repositories/{source.repoURL}/appdetails | GetAppDetails returns application details by given path
*RepositoryServiceApi* | [**RepositoryServiceGetHelmCharts**](docs/RepositoryServiceApi.md#repositoryservicegethelmcharts) | **Get** /api/v1/repositories/{repo}/helmcharts | GetHelmCharts returns list of helm charts in the specified repository
*RepositoryServiceApi* | [**RepositoryServiceListApps**](docs/RepositoryServiceApi.md#repositoryservicelistapps) | **Get** /api/v1/repositories/{repo}/apps | ListApps returns list of apps in the repo
*RepositoryServiceApi* | [**RepositoryServiceListRefs**](docs/RepositoryServiceApi.md#repositoryservicelistrefs) | **Get** /api/v1/repositories/{repo}/refs | 
*RepositoryServiceApi* | [**RepositoryServiceListRepositories**](docs/RepositoryServiceApi.md#repositoryservicelistrepositories) | **Get** /api/v1/repositories | ListRepositories gets a list of all configured repositories
*RepositoryServiceApi* | [**RepositoryServiceUpdateRepository**](docs/RepositoryServiceApi.md#repositoryserviceupdaterepository) | **Put** /api/v1/repositories/{repo.repo} | UpdateRepository updates a repository configuration
*RepositoryServiceApi* | [**RepositoryServiceValidateAccess**](docs/RepositoryServiceApi.md#repositoryservicevalidateaccess) | **Post** /api/v1/repositories/{repo}/validate | ValidateAccess validates access to a repository with given parameters
*SessionServiceApi* | [**SessionServiceCreate**](docs/SessionServiceApi.md#sessionservicecreate) | **Post** /api/v1/session | Create a new JWT for authentication and set a cookie if using HTTP
*SessionServiceApi* | [**SessionServiceDelete**](docs/SessionServiceApi.md#sessionservicedelete) | **Delete** /api/v1/session | Delete an existing JWT cookie if using HTTP
*SessionServiceApi* | [**SessionServiceGetUserInfo**](docs/SessionServiceApi.md#sessionservicegetuserinfo) | **Get** /api/v1/session/userinfo | Get the current user&#39;s info
*SettingsServiceApi* | [**SettingsServiceGet**](docs/SettingsServiceApi.md#settingsserviceget) | **Get** /api/v1/settings | Get returns Argo CD settings
*VersionServiceApi* | [**VersionServiceVersion**](docs/VersionServiceApi.md#versionserviceversion) | **Get** /api/version | Version returns version information of the API server


## Documentation For Models

 - [AccountAccount](docs/AccountAccount.md)
 - [AccountAccountsList](docs/AccountAccountsList.md)
 - [AccountCanIResponse](docs/AccountCanIResponse.md)
 - [AccountCreateTokenRequest](docs/AccountCreateTokenRequest.md)
 - [AccountCreateTokenResponse](docs/AccountCreateTokenResponse.md)
 - [AccountEmptyResponse](docs/AccountEmptyResponse.md)
 - [AccountToken](docs/AccountToken.md)
 - [AccountUpdatePasswordRequest](docs/AccountUpdatePasswordRequest.md)
 - [AccountUpdatePasswordResponse](docs/AccountUpdatePasswordResponse.md)
 - [ApplicationApplicationPatchRequest](docs/ApplicationApplicationPatchRequest.md)
 - [ApplicationApplicationResourceResponse](docs/ApplicationApplicationResourceResponse.md)
 - [ApplicationApplicationResponse](docs/ApplicationApplicationResponse.md)
 - [ApplicationApplicationRollbackRequest](docs/ApplicationApplicationRollbackRequest.md)
 - [ApplicationApplicationSyncRequest](docs/ApplicationApplicationSyncRequest.md)
 - [ApplicationApplicationSyncWindow](docs/ApplicationApplicationSyncWindow.md)
 - [ApplicationApplicationSyncWindowsResponse](docs/ApplicationApplicationSyncWindowsResponse.md)
 - [ApplicationLogEntry](docs/ApplicationLogEntry.md)
 - [ApplicationManagedResourcesResponse](docs/ApplicationManagedResourcesResponse.md)
 - [ApplicationOperationTerminateResponse](docs/ApplicationOperationTerminateResponse.md)
 - [ApplicationResourceActionsListResponse](docs/ApplicationResourceActionsListResponse.md)
 - [ApplicationSyncOptions](docs/ApplicationSyncOptions.md)
 - [Applicationv1alpha1EnvEntry](docs/Applicationv1alpha1EnvEntry.md)
 - [ClusterClusterResponse](docs/ClusterClusterResponse.md)
 - [ClusterConnector](docs/ClusterConnector.md)
 - [ClusterDexConfig](docs/ClusterDexConfig.md)
 - [ClusterGoogleAnalyticsConfig](docs/ClusterGoogleAnalyticsConfig.md)
 - [ClusterHelp](docs/ClusterHelp.md)
 - [ClusterOidcConfig](docs/ClusterOidcConfig.md)
 - [ClusterPlugin](docs/ClusterPlugin.md)
 - [ClusterSettings](docs/ClusterSettings.md)
 - [GpgkeyGnuPgPublicKeyCreateResponse](docs/GpgkeyGnuPgPublicKeyCreateResponse.md)
 - [GpgkeyGnuPgPublicKeyResponse](docs/GpgkeyGnuPgPublicKeyResponse.md)
 - [OidcClaim](docs/OidcClaim.md)
 - [ProjectDetailedProjectsResponse](docs/ProjectDetailedProjectsResponse.md)
 - [ProjectEmptyResponse](docs/ProjectEmptyResponse.md)
 - [ProjectGlobalProjectsResponse](docs/ProjectGlobalProjectsResponse.md)
 - [ProjectProjectCreateRequest](docs/ProjectProjectCreateRequest.md)
 - [ProjectProjectTokenCreateRequest](docs/ProjectProjectTokenCreateRequest.md)
 - [ProjectProjectTokenResponse](docs/ProjectProjectTokenResponse.md)
 - [ProjectProjectUpdateRequest](docs/ProjectProjectUpdateRequest.md)
 - [ProjectSyncWindowsResponse](docs/ProjectSyncWindowsResponse.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [RepocredsRepoCredsResponse](docs/RepocredsRepoCredsResponse.md)
 - [RepositoryAppInfo](docs/RepositoryAppInfo.md)
 - [RepositoryDirectoryAppSpec](docs/RepositoryDirectoryAppSpec.md)
 - [RepositoryHelmAppSpec](docs/RepositoryHelmAppSpec.md)
 - [RepositoryHelmChart](docs/RepositoryHelmChart.md)
 - [RepositoryHelmChartsResponse](docs/RepositoryHelmChartsResponse.md)
 - [RepositoryKsonnetAppSpec](docs/RepositoryKsonnetAppSpec.md)
 - [RepositoryKsonnetEnvironment](docs/RepositoryKsonnetEnvironment.md)
 - [RepositoryKsonnetEnvironmentDestination](docs/RepositoryKsonnetEnvironmentDestination.md)
 - [RepositoryKustomizeAppSpec](docs/RepositoryKustomizeAppSpec.md)
 - [RepositoryManifestResponse](docs/RepositoryManifestResponse.md)
 - [RepositoryRefs](docs/RepositoryRefs.md)
 - [RepositoryRepoAppDetailsQuery](docs/RepositoryRepoAppDetailsQuery.md)
 - [RepositoryRepoAppDetailsResponse](docs/RepositoryRepoAppDetailsResponse.md)
 - [RepositoryRepoAppsResponse](docs/RepositoryRepoAppsResponse.md)
 - [RepositoryRepoResponse](docs/RepositoryRepoResponse.md)
 - [RuntimeError](docs/RuntimeError.md)
 - [RuntimeStreamError](docs/RuntimeStreamError.md)
 - [SessionGetUserInfoResponse](docs/SessionGetUserInfoResponse.md)
 - [SessionSessionCreateRequest](docs/SessionSessionCreateRequest.md)
 - [SessionSessionResponse](docs/SessionSessionResponse.md)
 - [StreamResultOfApplicationLogEntry](docs/StreamResultOfApplicationLogEntry.md)
 - [StreamResultOfV1alpha1ApplicationTree](docs/StreamResultOfV1alpha1ApplicationTree.md)
 - [StreamResultOfV1alpha1ApplicationWatchEvent](docs/StreamResultOfV1alpha1ApplicationWatchEvent.md)
 - [V1Event](docs/V1Event.md)
 - [V1EventList](docs/V1EventList.md)
 - [V1EventSeries](docs/V1EventSeries.md)
 - [V1EventSource](docs/V1EventSource.md)
 - [V1FieldsV1](docs/V1FieldsV1.md)
 - [V1GroupKind](docs/V1GroupKind.md)
 - [V1ListMeta](docs/V1ListMeta.md)
 - [V1LoadBalancerIngress](docs/V1LoadBalancerIngress.md)
 - [V1ManagedFieldsEntry](docs/V1ManagedFieldsEntry.md)
 - [V1MicroTime](docs/V1MicroTime.md)
 - [V1NodeSystemInfo](docs/V1NodeSystemInfo.md)
 - [V1ObjectMeta](docs/V1ObjectMeta.md)
 - [V1ObjectReference](docs/V1ObjectReference.md)
 - [V1OwnerReference](docs/V1OwnerReference.md)
 - [V1PortStatus](docs/V1PortStatus.md)
 - [V1Time](docs/V1Time.md)
 - [V1alpha1AppProject](docs/V1alpha1AppProject.md)
 - [V1alpha1AppProjectList](docs/V1alpha1AppProjectList.md)
 - [V1alpha1AppProjectSpec](docs/V1alpha1AppProjectSpec.md)
 - [V1alpha1AppProjectStatus](docs/V1alpha1AppProjectStatus.md)
 - [V1alpha1Application](docs/V1alpha1Application.md)
 - [V1alpha1ApplicationCondition](docs/V1alpha1ApplicationCondition.md)
 - [V1alpha1ApplicationDestination](docs/V1alpha1ApplicationDestination.md)
 - [V1alpha1ApplicationList](docs/V1alpha1ApplicationList.md)
 - [V1alpha1ApplicationSource](docs/V1alpha1ApplicationSource.md)
 - [V1alpha1ApplicationSourceDirectory](docs/V1alpha1ApplicationSourceDirectory.md)
 - [V1alpha1ApplicationSourceHelm](docs/V1alpha1ApplicationSourceHelm.md)
 - [V1alpha1ApplicationSourceJsonnet](docs/V1alpha1ApplicationSourceJsonnet.md)
 - [V1alpha1ApplicationSourceKsonnet](docs/V1alpha1ApplicationSourceKsonnet.md)
 - [V1alpha1ApplicationSourceKustomize](docs/V1alpha1ApplicationSourceKustomize.md)
 - [V1alpha1ApplicationSourcePlugin](docs/V1alpha1ApplicationSourcePlugin.md)
 - [V1alpha1ApplicationSpec](docs/V1alpha1ApplicationSpec.md)
 - [V1alpha1ApplicationStatus](docs/V1alpha1ApplicationStatus.md)
 - [V1alpha1ApplicationSummary](docs/V1alpha1ApplicationSummary.md)
 - [V1alpha1ApplicationTree](docs/V1alpha1ApplicationTree.md)
 - [V1alpha1ApplicationWatchEvent](docs/V1alpha1ApplicationWatchEvent.md)
 - [V1alpha1AwsAuthConfig](docs/V1alpha1AwsAuthConfig.md)
 - [V1alpha1Backoff](docs/V1alpha1Backoff.md)
 - [V1alpha1Cluster](docs/V1alpha1Cluster.md)
 - [V1alpha1ClusterCacheInfo](docs/V1alpha1ClusterCacheInfo.md)
 - [V1alpha1ClusterConfig](docs/V1alpha1ClusterConfig.md)
 - [V1alpha1ClusterInfo](docs/V1alpha1ClusterInfo.md)
 - [V1alpha1ClusterList](docs/V1alpha1ClusterList.md)
 - [V1alpha1Command](docs/V1alpha1Command.md)
 - [V1alpha1ComparedTo](docs/V1alpha1ComparedTo.md)
 - [V1alpha1ConfigManagementPlugin](docs/V1alpha1ConfigManagementPlugin.md)
 - [V1alpha1ConnectionState](docs/V1alpha1ConnectionState.md)
 - [V1alpha1ExecProviderConfig](docs/V1alpha1ExecProviderConfig.md)
 - [V1alpha1GnuPgPublicKey](docs/V1alpha1GnuPgPublicKey.md)
 - [V1alpha1GnuPgPublicKeyList](docs/V1alpha1GnuPgPublicKeyList.md)
 - [V1alpha1HealthStatus](docs/V1alpha1HealthStatus.md)
 - [V1alpha1HelmFileParameter](docs/V1alpha1HelmFileParameter.md)
 - [V1alpha1HelmParameter](docs/V1alpha1HelmParameter.md)
 - [V1alpha1HostInfo](docs/V1alpha1HostInfo.md)
 - [V1alpha1HostResourceInfo](docs/V1alpha1HostResourceInfo.md)
 - [V1alpha1Info](docs/V1alpha1Info.md)
 - [V1alpha1InfoItem](docs/V1alpha1InfoItem.md)
 - [V1alpha1JsonnetVar](docs/V1alpha1JsonnetVar.md)
 - [V1alpha1JwtToken](docs/V1alpha1JwtToken.md)
 - [V1alpha1JwtTokens](docs/V1alpha1JwtTokens.md)
 - [V1alpha1KnownTypeField](docs/V1alpha1KnownTypeField.md)
 - [V1alpha1KsonnetParameter](docs/V1alpha1KsonnetParameter.md)
 - [V1alpha1KustomizeOptions](docs/V1alpha1KustomizeOptions.md)
 - [V1alpha1Operation](docs/V1alpha1Operation.md)
 - [V1alpha1OperationInitiator](docs/V1alpha1OperationInitiator.md)
 - [V1alpha1OperationState](docs/V1alpha1OperationState.md)
 - [V1alpha1OrphanedResourceKey](docs/V1alpha1OrphanedResourceKey.md)
 - [V1alpha1OrphanedResourcesMonitorSettings](docs/V1alpha1OrphanedResourcesMonitorSettings.md)
 - [V1alpha1OverrideIgnoreDiff](docs/V1alpha1OverrideIgnoreDiff.md)
 - [V1alpha1ProjectRole](docs/V1alpha1ProjectRole.md)
 - [V1alpha1RepoCreds](docs/V1alpha1RepoCreds.md)
 - [V1alpha1RepoCredsList](docs/V1alpha1RepoCredsList.md)
 - [V1alpha1Repository](docs/V1alpha1Repository.md)
 - [V1alpha1RepositoryCertificate](docs/V1alpha1RepositoryCertificate.md)
 - [V1alpha1RepositoryCertificateList](docs/V1alpha1RepositoryCertificateList.md)
 - [V1alpha1RepositoryList](docs/V1alpha1RepositoryList.md)
 - [V1alpha1ResourceAction](docs/V1alpha1ResourceAction.md)
 - [V1alpha1ResourceActionParam](docs/V1alpha1ResourceActionParam.md)
 - [V1alpha1ResourceDiff](docs/V1alpha1ResourceDiff.md)
 - [V1alpha1ResourceIgnoreDifferences](docs/V1alpha1ResourceIgnoreDifferences.md)
 - [V1alpha1ResourceNetworkingInfo](docs/V1alpha1ResourceNetworkingInfo.md)
 - [V1alpha1ResourceNode](docs/V1alpha1ResourceNode.md)
 - [V1alpha1ResourceOverride](docs/V1alpha1ResourceOverride.md)
 - [V1alpha1ResourceRef](docs/V1alpha1ResourceRef.md)
 - [V1alpha1ResourceResult](docs/V1alpha1ResourceResult.md)
 - [V1alpha1ResourceStatus](docs/V1alpha1ResourceStatus.md)
 - [V1alpha1RetryStrategy](docs/V1alpha1RetryStrategy.md)
 - [V1alpha1RevisionHistory](docs/V1alpha1RevisionHistory.md)
 - [V1alpha1RevisionMetadata](docs/V1alpha1RevisionMetadata.md)
 - [V1alpha1SignatureKey](docs/V1alpha1SignatureKey.md)
 - [V1alpha1SyncOperation](docs/V1alpha1SyncOperation.md)
 - [V1alpha1SyncOperationResource](docs/V1alpha1SyncOperationResource.md)
 - [V1alpha1SyncOperationResult](docs/V1alpha1SyncOperationResult.md)
 - [V1alpha1SyncPolicy](docs/V1alpha1SyncPolicy.md)
 - [V1alpha1SyncPolicyAutomated](docs/V1alpha1SyncPolicyAutomated.md)
 - [V1alpha1SyncStatus](docs/V1alpha1SyncStatus.md)
 - [V1alpha1SyncStrategy](docs/V1alpha1SyncStrategy.md)
 - [V1alpha1SyncStrategyApply](docs/V1alpha1SyncStrategyApply.md)
 - [V1alpha1SyncStrategyHook](docs/V1alpha1SyncStrategyHook.md)
 - [V1alpha1SyncWindow](docs/V1alpha1SyncWindow.md)
 - [V1alpha1TlsClientConfig](docs/V1alpha1TlsClientConfig.md)
 - [VersionVersionMessage](docs/VersionVersionMessage.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author

Alexander Morogov

