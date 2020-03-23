package constant

//NxLocalizedErrorCode
const (
	LocerrInternalError                                = -128
	LocerrGenericUnknown                               = -1
	LocerrGenericOk                                    = 0
	LocerrGenericNotSet                                = 1
	LocerrGenericNotFound                              = 2
	LocerrGenericAlreadyExists                         = 3
	LocerrGenericInvalidPath                           = 4
	LocerrGenericAccessDenied                          = 5
	LocerrGenericOutOfMemory                           = 6
	LocerrGenericNotInitialized                        = 7
	LocerrGenericInvalidParameters                     = 8
	LocerrGenericEmptyParameters                       = 9
	LocerrGenericInternalError                         = 10
	LocerrGenericCorruptData                           = 11
	LocerrGenericMemoryInconsistency                   = 12
	LocerrGenericInvisibleOwnerAbort                   = 13
	LocerrGenericProhibitValidate                      = 14
	LocerrGenericAborted                               = 15
	LocerrGenericConnectionLost                        = 16
	LocerrGenericUnsupportedInProductVersion           = 17
	LocerrGenericRestConnectionFailure                 = 18
	LocerrGenericMemoryLimitReached                    = 19
	LocerrHTTP400                                      = 400
	LocerrHTTP401                                      = 401
	LocerrHTTP402                                      = 402
	LocerrHTTP403                                      = 403
	LocerrHTTP404                                      = 404
	LocerrHTTP405                                      = 405
	LocerrHTTP406                                      = 406
	LocerrHTTP407                                      = 407
	LocerrHTTP408                                      = 408
	LocerrHTTP409                                      = 409
	LocerrHTTP410                                      = 410
	LocerrHTTP411                                      = 411
	LocerrHTTP412                                      = 412
	LocerrHTTP413                                      = 413
	LocerrHTTP414                                      = 414
	LocerrHTTP415                                      = 415
	LocerrHTTP416                                      = 416
	LocerrHTTP417                                      = 417
	LocerrHTTP422                                      = 422
	LocerrHTTP429                                      = 429
	LocerrHTTP500                                      = 500
	LocerrHTTP501                                      = 501
	LocerrHTTP502                                      = 502
	LocerrHTTP503                                      = 503
	LocerrHTTP504                                      = 504
	LocerrHTTP505                                      = 505
	LocerrHTTP509                                      = 509
	LocerrHTTPCouldNotResolveHost                      = 700
	LocerrAppAlreadyExists                             = 1000
	LocerrAppInvalidName                               = 1001
	LocerrAppAlreadyOpen                               = 1002
	LocerrAppNotFound                                  = 1003
	LocerrAppImportFailed                              = 1004
	LocerrAppSaveFailed                                = 1005
	LocerrAppCreateFailed                              = 1006
	LocerrAppInvalid                                   = 1007
	LocerrAppConnectFailed                             = 1008
	LocerrAppAlreadyOpenInDifferentMode                = 1009
	LocerrAppMigrationCouldNotContactMigrationService  = 1010
	LocerrAppMigrationCouldNotStartMigration           = 1011
	LocerrAppMigrationFailure                          = 1012
	LocerrAppScriptMissing                             = 1013
	LocerrAppExportFailed                              = 1014
	LocerrConnectionAlreadyExists                      = 2000
	LocerrConnectionNotFound                           = 2001
	LocerrConnectionFailedToLoad                       = 2002
	LocerrConnectionFailedToImport                     = 2003
	LocerrConnectionNameIsInvalid                      = 2004
	LocerrConnectorNoFileStreamingSupport              = 2300
	LocerrConnectorFilesizeExceededBufferSize          = 2301
	LocerrFileAccessDenied                             = 3000
	LocerrFileNameInvalid                              = 3001
	LocerrFileCorrupt                                  = 3002
	LocerrFileNotFound                                 = 3003
	LocerrFileFormatUnsupported                        = 3004
	LocerrFileOpenedInUnsupportedMode                  = 3005
	LocerrFileTableNotFound                            = 3006
	LocerrUserAccessDenied                             = 4000
	LocerrUserImpersonationFailed                      = 4001
	LocerrServerOutOfSessionAndUserCals                = 5000
	LocerrServerOutOfSessionCals                       = 5001
	LocerrServerOutOfUsageCals                         = 5002
	LocerrServerOutOfCals                              = 5003
	LocerrServerOutOfNamedCals                         = 5004
	LocerrServerOffDuty                                = 5005
	LocerrServerBusy                                   = 5006
	LocerrServerLicenseExpired                         = 5007
	LocerrServerAjaxDisabled                           = 5008
	LocerrHcInvalidObject                              = 6000
	LocerrHcResultTooLarge                             = 6001
	LocerrHcInvalidObjectState                         = 6002
	LocerrHcModalObjectError                           = 6003
	LocerrCalcInvalidDef                               = 7000
	LocerrCalcNotInLib                                 = 7001
	LocerrCalcHeapError                                = 7002
	LocerrCalcTooLarge                                 = 7003
	LocerrCalcTimeout                                  = 7004
	LocerrCalcEvalConditionFailed                      = 7005
	LocerrCalcMixedLinkedAggregation                   = 7006
	LocerrCalcMissingLinked                            = 7007
	LocerrCalcInvalidColSort                           = 7008
	LocerrCalcPagesTooLarge                            = 7009
	LocerrCalcSemanticFieldNotAllowed                  = 7010
	LocerrCalcValidationStateInvalid                   = 7011
	LocerrCalcPivotDimensionsAlreadyExists             = 7012
	LocerrCalcMissingLinkedField                       = 7013
	LocerrCalcNotCalculated                            = 7014
	LocerrLayoutExtendsInvalidID                       = 8000
	LocerrLayoutLinkedObjectNotFound                   = 8001
	LocerrLayoutLinkedObjectInvalid                    = 8002
	LocerrPersistenceWriteFailed                       = 9000
	LocerrPersistenceReadFailed                        = 9001
	LocerrPersistenceDeleteFailed                      = 9002
	LocerrPersistenceNotFound                          = 9003
	LocerrPersistenceUnsupportedVersion                = 9004
	LocerrPersistenceMigrationFailedReadOnly           = 9005
	LocerrPersistenceMigrationCancelled                = 9006
	LocerrPersistenceMigrationBackupFailed             = 9007
	LocerrPersistenceDiskFull                          = 9008
	LocerrPersistenceNotSupportedForSessionApp         = 9009
	LocerrPersistenceSyncSetChunkInvalidParameters     = 9510
	LocerrPersistenceSyncGetChunkInvalidParameters     = 9511
	LocerrScriptDatasourceAccessDenied                 = 10000
	LocerrReloadInProgress                             = 11000
	LocerrReloadTableXNotFound                         = 11001
	LocerrReloadUnknownStatement                       = 11002
	LocerrReloadExpectedSomethingFoundUnknown          = 11003
	LocerrReloadExpectedNothingFoundUnknown            = 11004
	LocerrReloadExpectedOneOf1TokensFoundUnknown       = 11005
	LocerrReloadExpectedOneOf2TokensFoundUnknown       = 11006
	LocerrReloadExpectedOneOf3TokensFoundUnknown       = 11007
	LocerrReloadExpectedOneOf4TokensFoundUnknown       = 11008
	LocerrReloadExpectedOneOf5TokensFoundUnknown       = 11009
	LocerrReloadExpectedOneOf6TokensFoundUnknown       = 11010
	LocerrReloadExpectedOneOf7TokensFoundUnknown       = 11011
	LocerrReloadExpectedOneOf8OrMoreTokensFoundUnknown = 11012
	LocerrReloadFieldXNotFound                         = 11013
	LocerrReloadMappingTableXNotFound                  = 11014
	LocerrReloadLibConnectionXNotFound                 = 11015
	LocerrReloadNameAlreadyTaken                       = 11016
	LocerrReloadWrongFileFormatDif                     = 11017
	LocerrReloadWrongFileFormatBiff                    = 11018
	LocerrReloadWrongFileFormatEncrypted               = 11019
	LocerrReloadOpenFileError                          = 11020
	LocerrReloadAutoGenerateCount                      = 11021
	LocerrReloadPeIllegalPrefixComb                    = 11022
	LocerrReloadMatchingControlStatementError          = 11023
	LocerrReloadMatchingLibpathXNotFound               = 11024
	LocerrReloadMatchingLibpathXInvalid                = 11025
	LocerrReloadMatchingLibpathXOutside                = 11026
	LocerrReloadNoQualifiedPathForFile                 = 11027
	LocerrReloadModeStatementOnlyForLibPaths           = 11028
	LocerrReloadInconsistentUseOfSemanticFields        = 11029
	LocerrReloadNoOpenDatabase                         = 11030
	LocerrReloadAggregationRequiredByGroupBy           = 11031
	LocerrReloadConnectMustUseLibPrefixInThisMode      = 11032
	LocerrReloadOdbcConnectFailed                      = 11033
	LocerrReloadOledbConnectFailed                     = 11034
	LocerrReloadCustomConnectFailed                    = 11035
	LocerrReloadOdbcReadFailed                         = 11036
	LocerrReloadOledbReadFailed                        = 11037
	LocerrReloadCustomReadFailed                       = 11038
	LocerrReloadBinaryLoadProhibited                   = 11039
	LocerrReloadConnectorStartFailed                   = 11040
	LocerrReloadConnectorNotResponding                 = 11041
	LocerrReloadConnectorReplyError                    = 11042
	LocerrReloadConnectorConnectError                  = 11043
	LocerrReloadConnectorNotFoundError                 = 11044
	LocerrReloadInputFieldWithDuplicateKeys            = 11045
	LocerrReloadConcatenateLoadNoPreviousTable         = 11046
	LocerrPersonalNewVersionAvailable                  = 12000
	LocerrPersonalVersionExpired                       = 12001
	LocerrPersonalSectionAccessDetected                = 12002
	LocerrPersonalAppDeletionFailed                    = 12003
	LocerrUserAuthenticationFailure                    = 12004
	LocerrExportOutOfMemory                            = 13000
	LocerrExportNoData                                 = 13001
	LocerrSyncInvalidOffset                            = 14000
	LocerrSearchTimeout                                = 15000
	LocerrDirectDiscoveryLinkedExpressionFail          = 16000
	LocerrDirectDiscoveryRowcountOverflow              = 16001
	LocerrDirectDiscoveryEmptyResult                   = 16002
	LocerrDirectDiscoveryDbConnectionFailed            = 16003
	LocerrDirectDiscoveryMeasureNotAllowed             = 16004
	LocerrDirectDiscoveryDetailNotAllowed              = 16005
	LocerrDirectDiscoveryNotSynthCircularAllowed       = 16006
	LocerrDirectDiscoveryOnlyOneDdTableAllowed         = 16007
	LocerrDirectDiscoveryDbAuthorizationFailed         = 16008
	LocerrSmartLoadTableNotFound                       = 17000
	LocerrSmartLoadTableDuplicated                     = 17001
	LocerrVariableNoName                               = 18000
	LocerrVariableDuplicateName                        = 18001
	LocerrVariableInconsistency                        = 18002
	LocerrMediaLibraryListFailed                       = 19000
	LocerrMediaLibraryContentFailed                    = 19001
	LocerrMediaBundlingFailed                          = 19002
	LocerrMediaUnbundlingFailed                        = 19003
	LocerrMediaLibraryNotFound                         = 19004
	LocerrFeatureDisabled                              = 20000
	LocerrJSONRPCInvalidRequest                        = 32600
	LocerrJSONRPCMethodNotFound                        = 32601
	LocerrJSONRPCInvalidParameters                     = 32602
	LocerrJSONRPCInternalError                         = 32603
	LocerrJSONRPCParseError                            = 32700
	LocerrMqSocketConnectFailure                       = 33000
	LocerrMqSocketOpenFailure                          = 33001
	LocerrMqProtocolNoRespone                          = 33002
	LocerrMqProtocolLibraryException                   = 33003
	LocerrMqProtocolConnectionClosed                   = 33004
	LocerrMqProtocolChannelClosed                      = 33005
	LocerrMqProtocolUnknownError                       = 33006
	LocerrMqProtocolInvalidStatus                      = 33007
	LocerrExtengineGrpcStatusOk                        = 22000
	LocerrExtengineGrpcStatusCancelled                 = 22001
	LocerrExtengineGrpcStatusUnknown                   = 22002
	LocerrExtengineGrpcStatusInvalidArgument           = 22003
	LocerrExtengineGrpcStatusDeadlineExceeded          = 22004
	LocerrExtengineGrpcStatusNotFound                  = 22005
	LocerrExtengineGrpcStatusAlreadyExists             = 22006
	LocerrExtengineGrpcStatusPermissionDenied          = 22007
	LocerrExtengineGrpcStatusResourceExhausted         = 22008
	LocerrExtengineGrpcStatusFailedPrecondition        = 22009
	LocerrExtengineGrpcStatusAborted                   = 22010
	LocerrExtengineGrpcStatusOutOfRange                = 22011
	LocerrExtengineGrpcStatusUnimplemented             = 22012
	LocerrExtengineGrpcStatusInternal                  = 22013
	LocerrExtengineGrpcStatusUnavailable               = 22014
	LocerrExtengineGrpcStatusDataLoss                  = 22015
	LocerrExtengineGrpcStatusUnauthenticated           = 22016
	LocerrLxwInvalidObj                                = 23001
	LocerrLxwInvalidFile                               = 23002
	LocerrLxwInvalidSheet                              = 23003
	LocerrLxwInvalidExportRange                        = 23004
	LocerrLxwError                                     = 23005
	LocerrLxwErrorMemoryMallocFailed                   = 23006
	LocerrLxwErrorCreatingXlsxFile                     = 23007
	LocerrLxwErrorCreatingTmpfile                      = 23008
	LocerrLxwErrorZipFileOperation                     = 23009
	LocerrLxwErrorZipFileAdd                           = 23010
	LocerrLxwErrorZipClose                             = 23011
	LocerrLxwErrorNullParameterIgnored                 = 23012
	LocerrLxwErrorMaxStringLengthExceeded              = 23013
	LocerrLxwError255StringLengthExceeded              = 23014
	LocerrLxwErrorSharedStringIndexNotFound            = 23015
	LocerrLxwErrorWorksheetIndexOutOfRange             = 23016
	LocerrLxwErrorWorksheetMaxNumberUrlsExceeded       = 23017
	LocerrCurlUnsupportedProtocol                      = 30000
	LocerrCurlCouldntResolveProxy                      = 30001
	LocerrCurlCouldntConnect                           = 30002
	LocerrCurlRemoteAccessDenied                       = 30003
	LocerrCurlFtpAcceptFailed                          = 30004
	LocerrCurlFtpAcceptTimeout                         = 30005
	LocerrCurlFtpCantGetHost                           = 30006
	LocerrCurlPartialFile                              = 30007
	LocerrCurlQuoteError                               = 30008
	LocerrCurlWriteError                               = 30009
	LocerrCurlUploadFailed                             = 30010
	LocerrCurlOutOfMemory                              = 30011
	LocerrCurlOperationTimedout                        = 30012
	LocerrCurlFtpCouldntUseRest                        = 30013
	LocerrCurlHTTPPostError                            = 30014
	LocerrCurlSslConnectError                          = 30015
	LocerrCurlFileCouldntReadFile                      = 30016
	LocerrCurlLdapCannotBind                           = 30017
	LocerrCurlLdapSearchFailed                         = 30018
	LocerrCurlTooManyRedirects                         = 30019
	LocerrCurlPeerFailedVerification                   = 30020
	LocerrCurlGotNothing                               = 30021
	LocerrCurlSslEngineNotfound                        = 30022
	LocerrCurlSslEngineSetfailed                       = 30023
	LocerrCurlSslCertproblem                           = 30024
	LocerrCurlSslCipher                                = 30025
	LocerrCurlSslCacert                                = 30026
	LocerrCurlBadContentEncoding                       = 30027
	LocerrCurlLdapInvalidURL                           = 30028
	LocerrCurlUseSslFailed                             = 30029
	LocerrCurlSslEngineInitfailed                      = 30030
	LocerrCurlLoginDenied                              = 30031
	LocerrCurlTftpNotfound                             = 30032
	LocerrCurlTftpIllegal                              = 30033
	LocerrCurlSSH                                      = 30034
)

//NxLocalizedWarningCode
const (
	LocwarnPersonalReloadRequired     = 0
	LocwarnPersonalVersionExpiresSoon = 1
	LocwarnExportDataTruncated        = 1000
	LocwarnCouldNotOpenAllObjects     = 2000
)
