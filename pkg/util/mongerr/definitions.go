package mongerr

// ErrOK -- code 0, ok
func ErrOK() Error {
	return Error{
		Message: "ok",
		Code:    0,
	}
}

// ErrInternalError -- code 1, internal error
func ErrInternalError() Error {
	return Error{
		Message: "internal error",
		Code:    1,
	}
}

// ErrBadValue -- code 2, bad value
func ErrBadValue() Error {
	return Error{
		Message: "bad value",
		Code:    2,
	}
}

// ErrOBSOLETE_DuplicateKey -- code 3, obsolete_ duplicate key
func ErrOBSOLETE_DuplicateKey() Error {
	return Error{
		Message: "obsolete_ duplicate key",
		Code:    3,
	}
}

// ErrNoSuchKey -- code 4, no such key
func ErrNoSuchKey() Error {
	return Error{
		Message: "no such key",
		Code:    4,
	}
}

// ErrGraphContainsCycle -- code 5, graph contains cycle
func ErrGraphContainsCycle() Error {
	return Error{
		Message: "graph contains cycle",
		Code:    5,
	}
}

// ErrHostUnreachable -- code 6, host unreachable
func ErrHostUnreachable() Error {
	return Error{
		Message: "host unreachable",
		Code:    6,
	}
}

// ErrHostNotFound -- code 7, host not found
func ErrHostNotFound() Error {
	return Error{
		Message: "host not found",
		Code:    7,
	}
}

// ErrUnknownError -- code 8, unknown error
func ErrUnknownError() Error {
	return Error{
		Message: "unknown error",
		Code:    8,
	}
}

// ErrFailedToParse -- code 9, failed to parse
func ErrFailedToParse() Error {
	return Error{
		Message: "failed to parse",
		Code:    9,
	}
}

// ErrCannotMutateObject -- code 10, cannot mutate object
func ErrCannotMutateObject() Error {
	return Error{
		Message: "cannot mutate object",
		Code:    10,
	}
}

// ErrUserNotFound -- code 11, user not found
func ErrUserNotFound() Error {
	return Error{
		Message: "user not found",
		Code:    11,
	}
}

// ErrUnsupportedFormat -- code 12, unsupported format
func ErrUnsupportedFormat() Error {
	return Error{
		Message: "unsupported format",
		Code:    12,
	}
}

// ErrUnauthorized -- code 13, unauthorized
func ErrUnauthorized() Error {
	return Error{
		Message: "unauthorized",
		Code:    13,
	}
}

// ErrTypeMismatch -- code 14, type mismatch
func ErrTypeMismatch() Error {
	return Error{
		Message: "type mismatch",
		Code:    14,
	}
}

// ErrOverflow -- code 15, overflow
func ErrOverflow() Error {
	return Error{
		Message: "overflow",
		Code:    15,
	}
}

// ErrInvalidLength -- code 16, invalid length
func ErrInvalidLength() Error {
	return Error{
		Message: "invalid length",
		Code:    16,
	}
}

// ErrProtocolError -- code 17, protocol error
func ErrProtocolError() Error {
	return Error{
		Message: "protocol error",
		Code:    17,
	}
}

// ErrAuthenticationFailed -- code 18, authentication failed
func ErrAuthenticationFailed() Error {
	return Error{
		Message: "authentication failed",
		Code:    18,
	}
}

// ErrCannotReuseObject -- code 19, cannot reuse object
func ErrCannotReuseObject() Error {
	return Error{
		Message: "cannot reuse object",
		Code:    19,
	}
}

// ErrIllegalOperation -- code 20, illegal operation
func ErrIllegalOperation() Error {
	return Error{
		Message: "illegal operation",
		Code:    20,
	}
}

// ErrEmptyArrayOperation -- code 21, empty array operation
func ErrEmptyArrayOperation() Error {
	return Error{
		Message: "empty array operation",
		Code:    21,
	}
}

// ErrInvalidBSON -- code 22, invalid bson
func ErrInvalidBSON() Error {
	return Error{
		Message: "invalid bson",
		Code:    22,
	}
}

// ErrAlreadyInitialized -- code 23, already initialized
func ErrAlreadyInitialized() Error {
	return Error{
		Message: "already initialized",
		Code:    23,
	}
}

// ErrLockTimeout -- code 24, lock timeout
func ErrLockTimeout() Error {
	return Error{
		Message: "lock timeout",
		Code:    24,
	}
}

// ErrRemoteValidationError -- code 25, remote validation error
func ErrRemoteValidationError() Error {
	return Error{
		Message: "remote validation error",
		Code:    25,
	}
}

// ErrNamespaceNotFound -- code 26, namespace not found
func ErrNamespaceNotFound() Error {
	return Error{
		Message: "namespace not found",
		Code:    26,
	}
}

// ErrIndexNotFound -- code 27, index not found
func ErrIndexNotFound() Error {
	return Error{
		Message: "index not found",
		Code:    27,
	}
}

// ErrPathNotViable -- code 28, path not viable
func ErrPathNotViable() Error {
	return Error{
		Message: "path not viable",
		Code:    28,
	}
}

// ErrNonExistentPath -- code 29, non existent path
func ErrNonExistentPath() Error {
	return Error{
		Message: "non existent path",
		Code:    29,
	}
}

// ErrInvalidPath -- code 30, invalid path
func ErrInvalidPath() Error {
	return Error{
		Message: "invalid path",
		Code:    30,
	}
}

// ErrRoleNotFound -- code 31, role not found
func ErrRoleNotFound() Error {
	return Error{
		Message: "role not found",
		Code:    31,
	}
}

// ErrRolesNotRelated -- code 32, roles not related
func ErrRolesNotRelated() Error {
	return Error{
		Message: "roles not related",
		Code:    32,
	}
}

// ErrPrivilegeNotFound -- code 33, privilege not found
func ErrPrivilegeNotFound() Error {
	return Error{
		Message: "privilege not found",
		Code:    33,
	}
}

// ErrCannotBackfillArray -- code 34, cannot backfill array
func ErrCannotBackfillArray() Error {
	return Error{
		Message: "cannot backfill array",
		Code:    34,
	}
}

// ErrUserModificationFailed -- code 35, user modification failed
func ErrUserModificationFailed() Error {
	return Error{
		Message: "user modification failed",
		Code:    35,
	}
}

// ErrRemoteChangeDetected -- code 36, remote change detected
func ErrRemoteChangeDetected() Error {
	return Error{
		Message: "remote change detected",
		Code:    36,
	}
}

// ErrFileRenameFailed -- code 37, file rename failed
func ErrFileRenameFailed() Error {
	return Error{
		Message: "file rename failed",
		Code:    37,
	}
}

// ErrFileNotOpen -- code 38, file not open
func ErrFileNotOpen() Error {
	return Error{
		Message: "file not open",
		Code:    38,
	}
}

// ErrFileStreamFailed -- code 39, file stream failed
func ErrFileStreamFailed() Error {
	return Error{
		Message: "file stream failed",
		Code:    39,
	}
}

// ErrConflictingUpdateOperators -- code 40, conflicting update operators
func ErrConflictingUpdateOperators() Error {
	return Error{
		Message: "conflicting update operators",
		Code:    40,
	}
}

// ErrFileAlreadyOpen -- code 41, file already open
func ErrFileAlreadyOpen() Error {
	return Error{
		Message: "file already open",
		Code:    41,
	}
}

// ErrLogWriteFailed -- code 42, log write failed
func ErrLogWriteFailed() Error {
	return Error{
		Message: "log write failed",
		Code:    42,
	}
}

// ErrCursorNotFound -- code 43, cursor not found
func ErrCursorNotFound() Error {
	return Error{
		Message: "cursor not found",
		Code:    43,
	}
}

// ErrUserDataInconsistent -- code 45, user data inconsistent
func ErrUserDataInconsistent() Error {
	return Error{
		Message: "user data inconsistent",
		Code:    45,
	}
}

// ErrLockBusy -- code 46, lock busy
func ErrLockBusy() Error {
	return Error{
		Message: "lock busy",
		Code:    46,
	}
}

// ErrNoMatchingDocument -- code 47, no matching document
func ErrNoMatchingDocument() Error {
	return Error{
		Message: "no matching document",
		Code:    47,
	}
}

// ErrNamespaceExists -- code 48, namespace exists
func ErrNamespaceExists() Error {
	return Error{
		Message: "namespace exists",
		Code:    48,
	}
}

// ErrInvalidRoleModification -- code 49, invalid role modification
func ErrInvalidRoleModification() Error {
	return Error{
		Message: "invalid role modification",
		Code:    49,
	}
}

// ErrExceededTimeLimit -- code 50, exceeded time limit
func ErrExceededTimeLimit() Error {
	return Error{
		Message: "exceeded time limit",
		Code:    50,
	}
}

// ErrManualInterventionRequired -- code 51, manual intervention required
func ErrManualInterventionRequired() Error {
	return Error{
		Message: "manual intervention required",
		Code:    51,
	}
}

// ErrDollarPrefixedFieldName -- code 52, dollar prefixed field name
func ErrDollarPrefixedFieldName() Error {
	return Error{
		Message: "dollar prefixed field name",
		Code:    52,
	}
}

// ErrInvalidIdField -- code 53, invalid id field
func ErrInvalidIdField() Error {
	return Error{
		Message: "invalid id field",
		Code:    53,
	}
}

// ErrNotSingleValueField -- code 54, not single value field
func ErrNotSingleValueField() Error {
	return Error{
		Message: "not single value field",
		Code:    54,
	}
}

// ErrInvalidDBRef -- code 55, invalid db ref
func ErrInvalidDBRef() Error {
	return Error{
		Message: "invalid db ref",
		Code:    55,
	}
}

// ErrEmptyFieldName -- code 56, empty field name
func ErrEmptyFieldName() Error {
	return Error{
		Message: "empty field name",
		Code:    56,
	}
}

// ErrDottedFieldName -- code 57, dotted field name
func ErrDottedFieldName() Error {
	return Error{
		Message: "dotted field name",
		Code:    57,
	}
}

// ErrRoleModificationFailed -- code 58, role modification failed
func ErrRoleModificationFailed() Error {
	return Error{
		Message: "role modification failed",
		Code:    58,
	}
}

// ErrCommandNotFound -- code 59, command not found
func ErrCommandNotFound() Error {
	return Error{
		Message: "command not found",
		Code:    59,
	}
}

// ErrOBSOLETE_DatabaseNotFound -- code 60, obsolete_ database not found
func ErrOBSOLETE_DatabaseNotFound() Error {
	return Error{
		Message: "obsolete_ database not found",
		Code:    60,
	}
}

// ErrShardKeyNotFound -- code 61, shard key not found
func ErrShardKeyNotFound() Error {
	return Error{
		Message: "shard key not found",
		Code:    61,
	}
}

// ErrOplogOperationUnsupported -- code 62, oplog operation unsupported
func ErrOplogOperationUnsupported() Error {
	return Error{
		Message: "oplog operation unsupported",
		Code:    62,
	}
}

// ErrStaleShardVersion -- code 63, stale shard version
func ErrStaleShardVersion() Error {
	return Error{
		Message: "stale shard version",
		Code:    63,
	}
}

// ErrWriteConcernFailed -- code 64, write concern failed
func ErrWriteConcernFailed() Error {
	return Error{
		Message: "write concern failed",
		Code:    64,
	}
}

// ErrMultipleErrorsOccurred -- code 65, multiple errors occurred
func ErrMultipleErrorsOccurred() Error {
	return Error{
		Message: "multiple errors occurred",
		Code:    65,
	}
}

// ErrImmutableField -- code 66, immutable field
func ErrImmutableField() Error {
	return Error{
		Message: "immutable field",
		Code:    66,
	}
}

// ErrCannotCreateIndex -- code 67, cannot create index
func ErrCannotCreateIndex() Error {
	return Error{
		Message: "cannot create index",
		Code:    67,
	}
}

// ErrIndexAlreadyExists -- code 68, index already exists
func ErrIndexAlreadyExists() Error {
	return Error{
		Message: "index already exists",
		Code:    68,
	}
}

// ErrAuthSchemaIncompatible -- code 69, auth schema incompatible
func ErrAuthSchemaIncompatible() Error {
	return Error{
		Message: "auth schema incompatible",
		Code:    69,
	}
}

// ErrShardNotFound -- code 70, shard not found
func ErrShardNotFound() Error {
	return Error{
		Message: "shard not found",
		Code:    70,
	}
}

// ErrReplicaSetNotFound -- code 71, replica set not found
func ErrReplicaSetNotFound() Error {
	return Error{
		Message: "replica set not found",
		Code:    71,
	}
}

// ErrInvalidOptions -- code 72, invalid options
func ErrInvalidOptions() Error {
	return Error{
		Message: "invalid options",
		Code:    72,
	}
}

// ErrInvalidNamespace -- code 73, invalid namespace
func ErrInvalidNamespace() Error {
	return Error{
		Message: "invalid namespace",
		Code:    73,
	}
}

// ErrNodeNotFound -- code 74, node not found
func ErrNodeNotFound() Error {
	return Error{
		Message: "node not found",
		Code:    74,
	}
}

// ErrWriteConcernLegacyOK -- code 75, write concern legacy ok
func ErrWriteConcernLegacyOK() Error {
	return Error{
		Message: "write concern legacy ok",
		Code:    75,
	}
}

// ErrNoReplicationEnabled -- code 76, no replication enabled
func ErrNoReplicationEnabled() Error {
	return Error{
		Message: "no replication enabled",
		Code:    76,
	}
}

// ErrOperationIncomplete -- code 77, operation incomplete
func ErrOperationIncomplete() Error {
	return Error{
		Message: "operation incomplete",
		Code:    77,
	}
}

// ErrCommandResultSchemaViolation -- code 78, command result schema violation
func ErrCommandResultSchemaViolation() Error {
	return Error{
		Message: "command result schema violation",
		Code:    78,
	}
}

// ErrUnknownReplWriteConcern -- code 79, unknown repl write concern
func ErrUnknownReplWriteConcern() Error {
	return Error{
		Message: "unknown repl write concern",
		Code:    79,
	}
}

// ErrRoleDataInconsistent -- code 80, role data inconsistent
func ErrRoleDataInconsistent() Error {
	return Error{
		Message: "role data inconsistent",
		Code:    80,
	}
}

// ErrNoMatchParseContext -- code 81, no match parse context
func ErrNoMatchParseContext() Error {
	return Error{
		Message: "no match parse context",
		Code:    81,
	}
}

// ErrNoProgressMade -- code 82, no progress made
func ErrNoProgressMade() Error {
	return Error{
		Message: "no progress made",
		Code:    82,
	}
}

// ErrRemoteResultsUnavailable -- code 83, remote results unavailable
func ErrRemoteResultsUnavailable() Error {
	return Error{
		Message: "remote results unavailable",
		Code:    83,
	}
}

// ErrDuplicateKeyValue -- code 84, duplicate key value
func ErrDuplicateKeyValue() Error {
	return Error{
		Message: "duplicate key value",
		Code:    84,
	}
}

// ErrIndexOptionsConflict -- code 85, index options conflict
func ErrIndexOptionsConflict() Error {
	return Error{
		Message: "index options conflict",
		Code:    85,
	}
}

// ErrIndexKeySpecsConflict -- code 86, index key specs conflict
func ErrIndexKeySpecsConflict() Error {
	return Error{
		Message: "index key specs conflict",
		Code:    86,
	}
}

// ErrCannotSplit -- code 87, cannot split
func ErrCannotSplit() Error {
	return Error{
		Message: "cannot split",
		Code:    87,
	}
}

// ErrSplitFailed_OBSOLETE -- code 88, split failed_obsolete
func ErrSplitFailed_OBSOLETE() Error {
	return Error{
		Message: "split failed_obsolete",
		Code:    88,
	}
}

// ErrNetworkTimeout -- code 89, network timeout
func ErrNetworkTimeout() Error {
	return Error{
		Message: "network timeout",
		Code:    89,
	}
}

// ErrCallbackCanceled -- code 90, callback canceled
func ErrCallbackCanceled() Error {
	return Error{
		Message: "callback canceled",
		Code:    90,
	}
}

// ErrShutdownInProgress -- code 91, shutdown in progress
func ErrShutdownInProgress() Error {
	return Error{
		Message: "shutdown in progress",
		Code:    91,
	}
}

// ErrSecondaryAheadOfPrimary -- code 92, secondary ahead of primary
func ErrSecondaryAheadOfPrimary() Error {
	return Error{
		Message: "secondary ahead of primary",
		Code:    92,
	}
}

// ErrInvalidReplicaSetConfig -- code 93, invalid replica set config
func ErrInvalidReplicaSetConfig() Error {
	return Error{
		Message: "invalid replica set config",
		Code:    93,
	}
}

// ErrNotYetInitialized -- code 94, not yet initialized
func ErrNotYetInitialized() Error {
	return Error{
		Message: "not yet initialized",
		Code:    94,
	}
}

// ErrNotSecondary -- code 95, not secondary
func ErrNotSecondary() Error {
	return Error{
		Message: "not secondary",
		Code:    95,
	}
}

// ErrOperationFailed -- code 96, operation failed
func ErrOperationFailed() Error {
	return Error{
		Message: "operation failed",
		Code:    96,
	}
}

// ErrNoProjectionFound -- code 97, no projection found
func ErrNoProjectionFound() Error {
	return Error{
		Message: "no projection found",
		Code:    97,
	}
}

// ErrDBPathInUse -- code 98, db path in use
func ErrDBPathInUse() Error {
	return Error{
		Message: "db path in use",
		Code:    98,
	}
}

// ErrCannotSatisfyWriteConcern -- code 100, cannot satisfy write concern
func ErrCannotSatisfyWriteConcern() Error {
	return Error{
		Message: "cannot satisfy write concern",
		Code:    100,
	}
}

// ErrOutdatedClient -- code 101, outdated client
func ErrOutdatedClient() Error {
	return Error{
		Message: "outdated client",
		Code:    101,
	}
}

// ErrIncompatibleAuditMetadata -- code 102, incompatible audit metadata
func ErrIncompatibleAuditMetadata() Error {
	return Error{
		Message: "incompatible audit metadata",
		Code:    102,
	}
}

// ErrNewReplicaSetConfigurationIncompatible -- code 103, new replica set configuration incompatible
func ErrNewReplicaSetConfigurationIncompatible() Error {
	return Error{
		Message: "new replica set configuration incompatible",
		Code:    103,
	}
}

// ErrNodeNotElectable -- code 104, node not electable
func ErrNodeNotElectable() Error {
	return Error{
		Message: "node not electable",
		Code:    104,
	}
}

// ErrIncompatibleShardingMetadata -- code 105, incompatible sharding metadata
func ErrIncompatibleShardingMetadata() Error {
	return Error{
		Message: "incompatible sharding metadata",
		Code:    105,
	}
}

// ErrDistributedClockSkewed -- code 106, distributed clock skewed
func ErrDistributedClockSkewed() Error {
	return Error{
		Message: "distributed clock skewed",
		Code:    106,
	}
}

// ErrLockFailed -- code 107, lock failed
func ErrLockFailed() Error {
	return Error{
		Message: "lock failed",
		Code:    107,
	}
}

// ErrInconsistentReplicaSetNames -- code 108, inconsistent replica set names
func ErrInconsistentReplicaSetNames() Error {
	return Error{
		Message: "inconsistent replica set names",
		Code:    108,
	}
}

// ErrConfigurationInProgress -- code 109, configuration in progress
func ErrConfigurationInProgress() Error {
	return Error{
		Message: "configuration in progress",
		Code:    109,
	}
}

// ErrCannotInitializeNodeWithData -- code 110, cannot initialize node with data
func ErrCannotInitializeNodeWithData() Error {
	return Error{
		Message: "cannot initialize node with data",
		Code:    110,
	}
}

// ErrNotExactValueField -- code 111, not exact value field
func ErrNotExactValueField() Error {
	return Error{
		Message: "not exact value field",
		Code:    111,
	}
}

// ErrWriteConflict -- code 112, write conflict
func ErrWriteConflict() Error {
	return Error{
		Message: "write conflict",
		Code:    112,
	}
}

// ErrInitialSyncFailure -- code 113, initial sync failure
func ErrInitialSyncFailure() Error {
	return Error{
		Message: "initial sync failure",
		Code:    113,
	}
}

// ErrInitialSyncOplogSourceMissing -- code 114, initial sync oplog source missing
func ErrInitialSyncOplogSourceMissing() Error {
	return Error{
		Message: "initial sync oplog source missing",
		Code:    114,
	}
}

// ErrCommandNotSupported -- code 115, command not supported
func ErrCommandNotSupported() Error {
	return Error{
		Message: "command not supported",
		Code:    115,
	}
}

// ErrDocTooLargeForCapped -- code 116, doc too large for capped
func ErrDocTooLargeForCapped() Error {
	return Error{
		Message: "doc too large for capped",
		Code:    116,
	}
}

// ErrConflictingOperationInProgress -- code 117, conflicting operation in progress
func ErrConflictingOperationInProgress() Error {
	return Error{
		Message: "conflicting operation in progress",
		Code:    117,
	}
}

// ErrNamespaceNotSharded -- code 118, namespace not sharded
func ErrNamespaceNotSharded() Error {
	return Error{
		Message: "namespace not sharded",
		Code:    118,
	}
}

// ErrInvalidSyncSource -- code 119, invalid sync source
func ErrInvalidSyncSource() Error {
	return Error{
		Message: "invalid sync source",
		Code:    119,
	}
}

// ErrOplogStartMissing -- code 120, oplog start missing
func ErrOplogStartMissing() Error {
	return Error{
		Message: "oplog start missing",
		Code:    120,
	}
}

// ErrDocumentValidationFailure -- code 121, document validation failure
func ErrDocumentValidationFailure() Error {
	return Error{
		Message: "document validation failure",
		Code:    121,
	}
}

// ErrOBSOLETE_ReadAfterOptimeTimeout -- code 122, obsolete_ read after optime timeout
func ErrOBSOLETE_ReadAfterOptimeTimeout() Error {
	return Error{
		Message: "obsolete_ read after optime timeout",
		Code:    122,
	}
}

// ErrNotAReplicaSet -- code 123, not a replica set
func ErrNotAReplicaSet() Error {
	return Error{
		Message: "not a replica set",
		Code:    123,
	}
}

// ErrIncompatibleElectionProtocol -- code 124, incompatible election protocol
func ErrIncompatibleElectionProtocol() Error {
	return Error{
		Message: "incompatible election protocol",
		Code:    124,
	}
}

// ErrCommandFailed -- code 125, command failed
func ErrCommandFailed() Error {
	return Error{
		Message: "command failed",
		Code:    125,
	}
}

// ErrRPCProtocolNegotiationFailed -- code 126, rpc protocol negotiation failed
func ErrRPCProtocolNegotiationFailed() Error {
	return Error{
		Message: "rpc protocol negotiation failed",
		Code:    126,
	}
}

// ErrUnrecoverableRollbackError -- code 127, unrecoverable rollback error
func ErrUnrecoverableRollbackError() Error {
	return Error{
		Message: "unrecoverable rollback error",
		Code:    127,
	}
}

// ErrLockNotFound -- code 128, lock not found
func ErrLockNotFound() Error {
	return Error{
		Message: "lock not found",
		Code:    128,
	}
}

// ErrLockStateChangeFailed -- code 129, lock state change failed
func ErrLockStateChangeFailed() Error {
	return Error{
		Message: "lock state change failed",
		Code:    129,
	}
}

// ErrSymbolNotFound -- code 130, symbol not found
func ErrSymbolNotFound() Error {
	return Error{
		Message: "symbol not found",
		Code:    130,
	}
}

// ErrRLPInitializationFailed -- code 131, rlp initialization failed
func ErrRLPInitializationFailed() Error {
	return Error{
		Message: "rlp initialization failed",
		Code:    131,
	}
}

// ErrOBSOLETE_ConfigServersInconsistent -- code 132, obsolete_ config servers inconsistent
func ErrOBSOLETE_ConfigServersInconsistent() Error {
	return Error{
		Message: "obsolete_ config servers inconsistent",
		Code:    132,
	}
}

// ErrFailedToSatisfyReadPreference -- code 133, failed to satisfy read preference
func ErrFailedToSatisfyReadPreference() Error {
	return Error{
		Message: "failed to satisfy read preference",
		Code:    133,
	}
}

// ErrReadConcernMajorityNotAvailableYet -- code 134, read concern majority not available yet
func ErrReadConcernMajorityNotAvailableYet() Error {
	return Error{
		Message: "read concern majority not available yet",
		Code:    134,
	}
}

// ErrStaleTerm -- code 135, stale term
func ErrStaleTerm() Error {
	return Error{
		Message: "stale term",
		Code:    135,
	}
}

// ErrCappedPositionLost -- code 136, capped position lost
func ErrCappedPositionLost() Error {
	return Error{
		Message: "capped position lost",
		Code:    136,
	}
}

// ErrIncompatibleShardingConfigVersion -- code 137, incompatible sharding config version
func ErrIncompatibleShardingConfigVersion() Error {
	return Error{
		Message: "incompatible sharding config version",
		Code:    137,
	}
}

// ErrRemoteOplogStale -- code 138, remote oplog stale
func ErrRemoteOplogStale() Error {
	return Error{
		Message: "remote oplog stale",
		Code:    138,
	}
}

// ErrJSInterpreterFailure -- code 139, js interpreter failure
func ErrJSInterpreterFailure() Error {
	return Error{
		Message: "js interpreter failure",
		Code:    139,
	}
}

// ErrInvalidSSLConfiguration -- code 140, invalid ssl configuration
func ErrInvalidSSLConfiguration() Error {
	return Error{
		Message: "invalid ssl configuration",
		Code:    140,
	}
}

// ErrSSLHandshakeFailed -- code 141, ssl handshake failed
func ErrSSLHandshakeFailed() Error {
	return Error{
		Message: "ssl handshake failed",
		Code:    141,
	}
}

// ErrJSUncatchableError -- code 142, js uncatchable error
func ErrJSUncatchableError() Error {
	return Error{
		Message: "js uncatchable error",
		Code:    142,
	}
}

// ErrCursorInUse -- code 143, cursor in use
func ErrCursorInUse() Error {
	return Error{
		Message: "cursor in use",
		Code:    143,
	}
}

// ErrIncompatibleCatalogManager -- code 144, incompatible catalog manager
func ErrIncompatibleCatalogManager() Error {
	return Error{
		Message: "incompatible catalog manager",
		Code:    144,
	}
}

// ErrPooledConnectionsDropped -- code 145, pooled connections dropped
func ErrPooledConnectionsDropped() Error {
	return Error{
		Message: "pooled connections dropped",
		Code:    145,
	}
}

// ErrExceededMemoryLimit -- code 146, exceeded memory limit
func ErrExceededMemoryLimit() Error {
	return Error{
		Message: "exceeded memory limit",
		Code:    146,
	}
}

// ErrZLibError -- code 147, z lib error
func ErrZLibError() Error {
	return Error{
		Message: "z lib error",
		Code:    147,
	}
}

// ErrReadConcernMajorityNotEnabled -- code 148, read concern majority not enabled
func ErrReadConcernMajorityNotEnabled() Error {
	return Error{
		Message: "read concern majority not enabled",
		Code:    148,
	}
}

// ErrNoConfigMaster -- code 149, no config master
func ErrNoConfigMaster() Error {
	return Error{
		Message: "no config master",
		Code:    149,
	}
}

// ErrStaleEpoch -- code 150, stale epoch
func ErrStaleEpoch() Error {
	return Error{
		Message: "stale epoch",
		Code:    150,
	}
}

// ErrOperationCannotBeBatched -- code 151, operation cannot be batched
func ErrOperationCannotBeBatched() Error {
	return Error{
		Message: "operation cannot be batched",
		Code:    151,
	}
}

// ErrOplogOutOfOrder -- code 152, oplog out of order
func ErrOplogOutOfOrder() Error {
	return Error{
		Message: "oplog out of order",
		Code:    152,
	}
}

// ErrChunkTooBig -- code 153, chunk too big
func ErrChunkTooBig() Error {
	return Error{
		Message: "chunk too big",
		Code:    153,
	}
}

// ErrInconsistentShardIdentity -- code 154, inconsistent shard identity
func ErrInconsistentShardIdentity() Error {
	return Error{
		Message: "inconsistent shard identity",
		Code:    154,
	}
}

// ErrCannotApplyOplogWhilePrimary -- code 155, cannot apply oplog while primary
func ErrCannotApplyOplogWhilePrimary() Error {
	return Error{
		Message: "cannot apply oplog while primary",
		Code:    155,
	}
}

// ErrNeedsDocumentMove -- code 156, needs document move
func ErrNeedsDocumentMove() Error {
	return Error{
		Message: "needs document move",
		Code:    156,
	}
}

// ErrCanRepairToDowngrade -- code 157, can repair to downgrade
func ErrCanRepairToDowngrade() Error {
	return Error{
		Message: "can repair to downgrade",
		Code:    157,
	}
}

// ErrMustUpgrade -- code 158, must upgrade
func ErrMustUpgrade() Error {
	return Error{
		Message: "must upgrade",
		Code:    158,
	}
}

// ErrDurationOverflow -- code 159, duration overflow
func ErrDurationOverflow() Error {
	return Error{
		Message: "duration overflow",
		Code:    159,
	}
}

// ErrMaxStalenessOutOfRange -- code 160, max staleness out of range
func ErrMaxStalenessOutOfRange() Error {
	return Error{
		Message: "max staleness out of range",
		Code:    160,
	}
}

// ErrIncompatibleCollationVersion -- code 161, incompatible collation version
func ErrIncompatibleCollationVersion() Error {
	return Error{
		Message: "incompatible collation version",
		Code:    161,
	}
}

// ErrCollectionIsEmpty -- code 162, collection is empty
func ErrCollectionIsEmpty() Error {
	return Error{
		Message: "collection is empty",
		Code:    162,
	}
}

// ErrZoneStillInUse -- code 163, zone still in use
func ErrZoneStillInUse() Error {
	return Error{
		Message: "zone still in use",
		Code:    163,
	}
}

// ErrInitialSyncActive -- code 164, initial sync active
func ErrInitialSyncActive() Error {
	return Error{
		Message: "initial sync active",
		Code:    164,
	}
}

// ErrViewDepthLimitExceeded -- code 165, view depth limit exceeded
func ErrViewDepthLimitExceeded() Error {
	return Error{
		Message: "view depth limit exceeded",
		Code:    165,
	}
}

// ErrCommandNotSupportedOnView -- code 166, command not supported on view
func ErrCommandNotSupportedOnView() Error {
	return Error{
		Message: "command not supported on view",
		Code:    166,
	}
}

// ErrOptionNotSupportedOnView -- code 167, option not supported on view
func ErrOptionNotSupportedOnView() Error {
	return Error{
		Message: "option not supported on view",
		Code:    167,
	}
}

// ErrInvalidPipelineOperator -- code 168, invalid pipeline operator
func ErrInvalidPipelineOperator() Error {
	return Error{
		Message: "invalid pipeline operator",
		Code:    168,
	}
}

// ErrCommandOnShardedViewNotSupportedOnMongod -- code 169, command on sharded view not supported on mongod
func ErrCommandOnShardedViewNotSupportedOnMongod() Error {
	return Error{
		Message: "command on sharded view not supported on mongod",
		Code:    169,
	}
}

// ErrTooManyMatchingDocuments -- code 170, too many matching documents
func ErrTooManyMatchingDocuments() Error {
	return Error{
		Message: "too many matching documents",
		Code:    170,
	}
}

// ErrCannotIndexParallelArrays -- code 171, cannot index parallel arrays
func ErrCannotIndexParallelArrays() Error {
	return Error{
		Message: "cannot index parallel arrays",
		Code:    171,
	}
}

// ErrTransportSessionClosed -- code 172, transport session closed
func ErrTransportSessionClosed() Error {
	return Error{
		Message: "transport session closed",
		Code:    172,
	}
}

// ErrTransportSessionNotFound -- code 173, transport session not found
func ErrTransportSessionNotFound() Error {
	return Error{
		Message: "transport session not found",
		Code:    173,
	}
}

// ErrTransportSessionUnknown -- code 174, transport session unknown
func ErrTransportSessionUnknown() Error {
	return Error{
		Message: "transport session unknown",
		Code:    174,
	}
}

// ErrQueryPlanKilled -- code 175, query plan killed
func ErrQueryPlanKilled() Error {
	return Error{
		Message: "query plan killed",
		Code:    175,
	}
}

// ErrFileOpenFailed -- code 176, file open failed
func ErrFileOpenFailed() Error {
	return Error{
		Message: "file open failed",
		Code:    176,
	}
}

// ErrZoneNotFound -- code 177, zone not found
func ErrZoneNotFound() Error {
	return Error{
		Message: "zone not found",
		Code:    177,
	}
}

// ErrRangeOverlapConflict -- code 178, range overlap conflict
func ErrRangeOverlapConflict() Error {
	return Error{
		Message: "range overlap conflict",
		Code:    178,
	}
}

// ErrWindowsPdhError -- code 179, windows pdh error
func ErrWindowsPdhError() Error {
	return Error{
		Message: "windows pdh error",
		Code:    179,
	}
}

// ErrBadPerfCounterPath -- code 180, bad perf counter path
func ErrBadPerfCounterPath() Error {
	return Error{
		Message: "bad perf counter path",
		Code:    180,
	}
}

// ErrAmbiguousIndexKeyPattern -- code 181, ambiguous index key pattern
func ErrAmbiguousIndexKeyPattern() Error {
	return Error{
		Message: "ambiguous index key pattern",
		Code:    181,
	}
}

// ErrInvalidViewDefinition -- code 182, invalid view definition
func ErrInvalidViewDefinition() Error {
	return Error{
		Message: "invalid view definition",
		Code:    182,
	}
}

// ErrClientMetadataMissingField -- code 183, client metadata missing field
func ErrClientMetadataMissingField() Error {
	return Error{
		Message: "client metadata missing field",
		Code:    183,
	}
}

// ErrClientMetadataAppNameTooLarge -- code 184, client metadata app name too large
func ErrClientMetadataAppNameTooLarge() Error {
	return Error{
		Message: "client metadata app name too large",
		Code:    184,
	}
}

// ErrClientMetadataDocumentTooLarge -- code 185, client metadata document too large
func ErrClientMetadataDocumentTooLarge() Error {
	return Error{
		Message: "client metadata document too large",
		Code:    185,
	}
}

// ErrClientMetadataCannotBeMutated -- code 186, client metadata cannot be mutated
func ErrClientMetadataCannotBeMutated() Error {
	return Error{
		Message: "client metadata cannot be mutated",
		Code:    186,
	}
}

// ErrLinearizableReadConcernError -- code 187, linearizable read concern error
func ErrLinearizableReadConcernError() Error {
	return Error{
		Message: "linearizable read concern error",
		Code:    187,
	}
}

// ErrIncompatibleServerVersion -- code 188, incompatible server version
func ErrIncompatibleServerVersion() Error {
	return Error{
		Message: "incompatible server version",
		Code:    188,
	}
}

// ErrPrimarySteppedDown -- code 189, primary stepped down
func ErrPrimarySteppedDown() Error {
	return Error{
		Message: "primary stepped down",
		Code:    189,
	}
}

// ErrMasterSlaveConnectionFailure -- code 190, master slave connection failure
func ErrMasterSlaveConnectionFailure() Error {
	return Error{
		Message: "master slave connection failure",
		Code:    190,
	}
}

// ErrOBSOLETE_BalancerLostDistributedLock -- code 191, obsolete_ balancer lost distributed lock
func ErrOBSOLETE_BalancerLostDistributedLock() Error {
	return Error{
		Message: "obsolete_ balancer lost distributed lock",
		Code:    191,
	}
}

// ErrFailPointEnabled -- code 192, fail point enabled
func ErrFailPointEnabled() Error {
	return Error{
		Message: "fail point enabled",
		Code:    192,
	}
}

// ErrNoShardingEnabled -- code 193, no sharding enabled
func ErrNoShardingEnabled() Error {
	return Error{
		Message: "no sharding enabled",
		Code:    193,
	}
}

// ErrBalancerInterrupted -- code 194, balancer interrupted
func ErrBalancerInterrupted() Error {
	return Error{
		Message: "balancer interrupted",
		Code:    194,
	}
}

// ErrViewPipelineMaxSizeExceeded -- code 195, view pipeline max size exceeded
func ErrViewPipelineMaxSizeExceeded() Error {
	return Error{
		Message: "view pipeline max size exceeded",
		Code:    195,
	}
}

// ErrInvalidIndexSpecificationOption -- code 197, invalid index specification option
func ErrInvalidIndexSpecificationOption() Error {
	return Error{
		Message: "invalid index specification option",
		Code:    197,
	}
}

// ErrOBSOLETE_ReceivedOpReplyMessage -- code 198, obsolete_ received op reply message
func ErrOBSOLETE_ReceivedOpReplyMessage() Error {
	return Error{
		Message: "obsolete_ received op reply message",
		Code:    198,
	}
}

// ErrReplicaSetMonitorRemoved -- code 199, replica set monitor removed
func ErrReplicaSetMonitorRemoved() Error {
	return Error{
		Message: "replica set monitor removed",
		Code:    199,
	}
}

// ErrChunkRangeCleanupPending -- code 200, chunk range cleanup pending
func ErrChunkRangeCleanupPending() Error {
	return Error{
		Message: "chunk range cleanup pending",
		Code:    200,
	}
}

// ErrCannotBuildIndexKeys -- code 201, cannot build index keys
func ErrCannotBuildIndexKeys() Error {
	return Error{
		Message: "cannot build index keys",
		Code:    201,
	}
}

// ErrNetworkInterfaceExceededTimeLimit -- code 202, network interface exceeded time limit
func ErrNetworkInterfaceExceededTimeLimit() Error {
	return Error{
		Message: "network interface exceeded time limit",
		Code:    202,
	}
}

// ErrShardingStateNotInitialized -- code 203, sharding state not initialized
func ErrShardingStateNotInitialized() Error {
	return Error{
		Message: "sharding state not initialized",
		Code:    203,
	}
}

// ErrTimeProofMismatch -- code 204, time proof mismatch
func ErrTimeProofMismatch() Error {
	return Error{
		Message: "time proof mismatch",
		Code:    204,
	}
}

// ErrClusterTimeFailsRateLimiter -- code 205, cluster time fails rate limiter
func ErrClusterTimeFailsRateLimiter() Error {
	return Error{
		Message: "cluster time fails rate limiter",
		Code:    205,
	}
}

// ErrNoSuchSession -- code 206, no such session
func ErrNoSuchSession() Error {
	return Error{
		Message: "no such session",
		Code:    206,
	}
}

// ErrInvalidUUID -- code 207, invalid uuid
func ErrInvalidUUID() Error {
	return Error{
		Message: "invalid uuid",
		Code:    207,
	}
}

// ErrTooManyLocks -- code 208, too many locks
func ErrTooManyLocks() Error {
	return Error{
		Message: "too many locks",
		Code:    208,
	}
}

// ErrStaleClusterTime -- code 209, stale cluster time
func ErrStaleClusterTime() Error {
	return Error{
		Message: "stale cluster time",
		Code:    209,
	}
}

// ErrCannotVerifyAndSignLogicalTime -- code 210, cannot verify and sign logical time
func ErrCannotVerifyAndSignLogicalTime() Error {
	return Error{
		Message: "cannot verify and sign logical time",
		Code:    210,
	}
}

// ErrKeyNotFound -- code 211, key not found
func ErrKeyNotFound() Error {
	return Error{
		Message: "key not found",
		Code:    211,
	}
}

// ErrIncompatibleRollbackAlgorithm -- code 212, incompatible rollback algorithm
func ErrIncompatibleRollbackAlgorithm() Error {
	return Error{
		Message: "incompatible rollback algorithm",
		Code:    212,
	}
}

// ErrDuplicateSession -- code 213, duplicate session
func ErrDuplicateSession() Error {
	return Error{
		Message: "duplicate session",
		Code:    213,
	}
}

// ErrAuthenticationRestrictionUnmet -- code 214, authentication restriction unmet
func ErrAuthenticationRestrictionUnmet() Error {
	return Error{
		Message: "authentication restriction unmet",
		Code:    214,
	}
}

// ErrDatabaseDropPending -- code 215, database drop pending
func ErrDatabaseDropPending() Error {
	return Error{
		Message: "database drop pending",
		Code:    215,
	}
}

// ErrElectionInProgress -- code 216, election in progress
func ErrElectionInProgress() Error {
	return Error{
		Message: "election in progress",
		Code:    216,
	}
}

// ErrIncompleteTransactionHistory -- code 217, incomplete transaction history
func ErrIncompleteTransactionHistory() Error {
	return Error{
		Message: "incomplete transaction history",
		Code:    217,
	}
}

// ErrUpdateOperationFailed -- code 218, update operation failed
func ErrUpdateOperationFailed() Error {
	return Error{
		Message: "update operation failed",
		Code:    218,
	}
}

// ErrFTDCPathNotSet -- code 219, ftdc path not set
func ErrFTDCPathNotSet() Error {
	return Error{
		Message: "ftdc path not set",
		Code:    219,
	}
}

// ErrFTDCPathAlreadySet -- code 220, ftdc path already set
func ErrFTDCPathAlreadySet() Error {
	return Error{
		Message: "ftdc path already set",
		Code:    220,
	}
}

// ErrIndexModified -- code 221, index modified
func ErrIndexModified() Error {
	return Error{
		Message: "index modified",
		Code:    221,
	}
}

// ErrCloseChangeStream -- code 222, close change stream
func ErrCloseChangeStream() Error {
	return Error{
		Message: "close change stream",
		Code:    222,
	}
}

// ErrIllegalOpMsgFlag -- code 223, illegal op msg flag
func ErrIllegalOpMsgFlag() Error {
	return Error{
		Message: "illegal op msg flag",
		Code:    223,
	}
}

// ErrQueryFeatureNotAllowed -- code 224, query feature not allowed
func ErrQueryFeatureNotAllowed() Error {
	return Error{
		Message: "query feature not allowed",
		Code:    224,
	}
}

// ErrTransactionTooOld -- code 225, transaction too old
func ErrTransactionTooOld() Error {
	return Error{
		Message: "transaction too old",
		Code:    225,
	}
}

// ErrAtomicityFailure -- code 226, atomicity failure
func ErrAtomicityFailure() Error {
	return Error{
		Message: "atomicity failure",
		Code:    226,
	}
}

// ErrCannotImplicitlyCreateCollection -- code 227, cannot implicitly create collection
func ErrCannotImplicitlyCreateCollection() Error {
	return Error{
		Message: "cannot implicitly create collection",
		Code:    227,
	}
}

// ErrSessionTransferIncomplete -- code 228, session transfer incomplete
func ErrSessionTransferIncomplete() Error {
	return Error{
		Message: "session transfer incomplete",
		Code:    228,
	}
}

// ErrMustDowngrade -- code 229, must downgrade
func ErrMustDowngrade() Error {
	return Error{
		Message: "must downgrade",
		Code:    229,
	}
}

// ErrDNSHostNotFound -- code 230, dns host not found
func ErrDNSHostNotFound() Error {
	return Error{
		Message: "dns host not found",
		Code:    230,
	}
}

// ErrDNSProtocolError -- code 231, dns protocol error
func ErrDNSProtocolError() Error {
	return Error{
		Message: "dns protocol error",
		Code:    231,
	}
}

// ErrMaxSubPipelineDepthExceeded -- code 232, max sub pipeline depth exceeded
func ErrMaxSubPipelineDepthExceeded() Error {
	return Error{
		Message: "max sub pipeline depth exceeded",
		Code:    232,
	}
}

// ErrTooManyDocumentSequences -- code 233, too many document sequences
func ErrTooManyDocumentSequences() Error {
	return Error{
		Message: "too many document sequences",
		Code:    233,
	}
}

// ErrRetryChangeStream -- code 234, retry change stream
func ErrRetryChangeStream() Error {
	return Error{
		Message: "retry change stream",
		Code:    234,
	}
}

// ErrInternalErrorNotSupported -- code 235, internal error not supported
func ErrInternalErrorNotSupported() Error {
	return Error{
		Message: "internal error not supported",
		Code:    235,
	}
}

// ErrForTestingErrorExtraInfo -- code 236, for testing error extra info
func ErrForTestingErrorExtraInfo() Error {
	return Error{
		Message: "for testing error extra info",
		Code:    236,
	}
}

// ErrCursorKilled -- code 237, cursor killed
func ErrCursorKilled() Error {
	return Error{
		Message: "cursor killed",
		Code:    237,
	}
}

// ErrNotImplemented -- code 238, not implemented
func ErrNotImplemented() Error {
	return Error{
		Message: "not implemented",
		Code:    238,
	}
}

// ErrSnapshotTooOld -- code 239, snapshot too old
func ErrSnapshotTooOld() Error {
	return Error{
		Message: "snapshot too old",
		Code:    239,
	}
}

// ErrDNSRecordTypeMismatch -- code 240, dns record type mismatch
func ErrDNSRecordTypeMismatch() Error {
	return Error{
		Message: "dns record type mismatch",
		Code:    240,
	}
}

// ErrConversionFailure -- code 241, conversion failure
func ErrConversionFailure() Error {
	return Error{
		Message: "conversion failure",
		Code:    241,
	}
}

// ErrCannotCreateCollection -- code 242, cannot create collection
func ErrCannotCreateCollection() Error {
	return Error{
		Message: "cannot create collection",
		Code:    242,
	}
}

// ErrIncompatibleWithUpgradedServer -- code 243, incompatible with upgraded server
func ErrIncompatibleWithUpgradedServer() Error {
	return Error{
		Message: "incompatible with upgraded server",
		Code:    243,
	}
}

// ErrNOT_YET_AVAILABLE_TransactionAborted -- code 244, not_yet_available_ transaction aborted
func ErrNOT_YET_AVAILABLE_TransactionAborted() Error {
	return Error{
		Message: "not_yet_available_ transaction aborted",
		Code:    244,
	}
}

// ErrBrokenPromise -- code 245, broken promise
func ErrBrokenPromise() Error {
	return Error{
		Message: "broken promise",
		Code:    245,
	}
}

// ErrSnapshotUnavailable -- code 246, snapshot unavailable
func ErrSnapshotUnavailable() Error {
	return Error{
		Message: "snapshot unavailable",
		Code:    246,
	}
}

// ErrProducerConsumerQueueBatchTooLarge -- code 247, producer consumer queue batch too large
func ErrProducerConsumerQueueBatchTooLarge() Error {
	return Error{
		Message: "producer consumer queue batch too large",
		Code:    247,
	}
}

// ErrProducerConsumerQueueEndClosed -- code 248, producer consumer queue end closed
func ErrProducerConsumerQueueEndClosed() Error {
	return Error{
		Message: "producer consumer queue end closed",
		Code:    248,
	}
}

// ErrStaleDbVersion -- code 249, stale db version
func ErrStaleDbVersion() Error {
	return Error{
		Message: "stale db version",
		Code:    249,
	}
}

// ErrStaleChunkHistory -- code 250, stale chunk history
func ErrStaleChunkHistory() Error {
	return Error{
		Message: "stale chunk history",
		Code:    250,
	}
}

// ErrNoSuchTransaction -- code 251, no such transaction
func ErrNoSuchTransaction() Error {
	return Error{
		Message: "no such transaction",
		Code:    251,
	}
}

// ErrReentrancyNotAllowed -- code 252, reentrancy not allowed
func ErrReentrancyNotAllowed() Error {
	return Error{
		Message: "reentrancy not allowed",
		Code:    252,
	}
}

// ErrFreeMonHttpInFlight -- code 253, free mon http in flight
func ErrFreeMonHttpInFlight() Error {
	return Error{
		Message: "free mon http in flight",
		Code:    253,
	}
}

// ErrFreeMonHttpTemporaryFailure -- code 254, free mon http temporary failure
func ErrFreeMonHttpTemporaryFailure() Error {
	return Error{
		Message: "free mon http temporary failure",
		Code:    254,
	}
}

// ErrFreeMonHttpPermanentFailure -- code 255, free mon http permanent failure
func ErrFreeMonHttpPermanentFailure() Error {
	return Error{
		Message: "free mon http permanent failure",
		Code:    255,
	}
}

// ErrTransactionCommitted -- code 256, transaction committed
func ErrTransactionCommitted() Error {
	return Error{
		Message: "transaction committed",
		Code:    256,
	}
}

// ErrTransactionTooLarge -- code 257, transaction too large
func ErrTransactionTooLarge() Error {
	return Error{
		Message: "transaction too large",
		Code:    257,
	}
}

// ErrUnknownFeatureCompatibilityVersion -- code 258, unknown feature compatibility version
func ErrUnknownFeatureCompatibilityVersion() Error {
	return Error{
		Message: "unknown feature compatibility version",
		Code:    258,
	}
}

// ErrSocketException -- code 9001, socket exception
func ErrSocketException() Error {
	return Error{
		Message: "socket exception",
		Code:    9001,
	}
}

// ErrOBSOLETE_RecvStaleConfig -- code 9996, obsolete_ recv stale config
func ErrOBSOLETE_RecvStaleConfig() Error {
	return Error{
		Message: "obsolete_ recv stale config",
		Code:    9996,
	}
}

// ErrNotMaster -- code 10107, not master
func ErrNotMaster() Error {
	return Error{
		Message: "not master",
		Code:    10107,
	}
}

// ErrCannotGrowDocumentInCappedNamespace -- code 10003, cannot grow document in capped namespace
func ErrCannotGrowDocumentInCappedNamespace() Error {
	return Error{
		Message: "cannot grow document in capped namespace",
		Code:    10003,
	}
}

// ErrBSONObjectTooLarge -- code 10334, bson object too large
func ErrBSONObjectTooLarge() Error {
	return Error{
		Message: "bson object too large",
		Code:    10334,
	}
}

// ErrDuplicateKey -- code 11000, duplicate key
func ErrDuplicateKey() Error {
	return Error{
		Message: "duplicate key",
		Code:    11000,
	}
}

// ErrInterruptedAtShutdown -- code 11600, interrupted at shutdown
func ErrInterruptedAtShutdown() Error {
	return Error{
		Message: "interrupted at shutdown",
		Code:    11600,
	}
}

// ErrInterrupted -- code 11601, interrupted
func ErrInterrupted() Error {
	return Error{
		Message: "interrupted",
		Code:    11601,
	}
}

// ErrInterruptedDueToReplStateChange -- code 11602, interrupted due to repl state change
func ErrInterruptedDueToReplStateChange() Error {
	return Error{
		Message: "interrupted due to repl state change",
		Code:    11602,
	}
}

// ErrOutOfDiskSpace -- code 14031, out of disk space
func ErrOutOfDiskSpace() Error {
	return Error{
		Message: "out of disk space",
		Code:    14031,
	}
}

// ErrKeyTooLong -- code 17280, key too long
func ErrKeyTooLong() Error {
	return Error{
		Message: "key too long",
		Code:    17280,
	}
}

// ErrBackgroundOperationInProgressForDatabase -- code 12586, background operation in progress for database
func ErrBackgroundOperationInProgressForDatabase() Error {
	return Error{
		Message: "background operation in progress for database",
		Code:    12586,
	}
}

// ErrBackgroundOperationInProgressForNamespace -- code 12587, background operation in progress for namespace
func ErrBackgroundOperationInProgressForNamespace() Error {
	return Error{
		Message: "background operation in progress for namespace",
		Code:    12587,
	}
}

// ErrNotMasterOrSecondary -- code 13436, not master or secondary
func ErrNotMasterOrSecondary() Error {
	return Error{
		Message: "not master or secondary",
		Code:    13436,
	}
}

// ErrNotMasterNoSlaveOk -- code 13435, not master no slave ok
func ErrNotMasterNoSlaveOk() Error {
	return Error{
		Message: "not master no slave ok",
		Code:    13435,
	}
}

// ErrShardKeyTooBig -- code 13334, shard key too big
func ErrShardKeyTooBig() Error {
	return Error{
		Message: "shard key too big",
		Code:    13334,
	}
}

// ErrStaleConfig -- code 13388, stale config
func ErrStaleConfig() Error {
	return Error{
		Message: "stale config",
		Code:    13388,
	}
}

// ErrDatabaseDifferCase -- code 13297, database differ case
func ErrDatabaseDifferCase() Error {
	return Error{
		Message: "database differ case",
		Code:    13297,
	}
}

// ErrOBSOLETE_PrepareConfigsFailed -- code 13104, obsolete_ prepare configs failed
func ErrOBSOLETE_PrepareConfigsFailed() Error {
	return Error{
		Message: "obsolete_ prepare configs failed",
		Code:    13104,
	}
}
