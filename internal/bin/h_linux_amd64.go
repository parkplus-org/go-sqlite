// Code generated by `gocc /tmp/go-generate-218625601/x.c -o internal/bin/h_linux_amd64.go -Itestdata/sqlite-amalgamation-3300100 -qbec-defines -qbec-enumconsts -qbec-import <none> -qbec-pkgname bin -qbec-structs -DHAVE_USLEEP -DLONGDOUBLE_TYPE=double -DSQLITE_DEFAULT_MEMSTATUS=0 -DSQLITE_DEFAULT_WAL_SYNCHRONOUS=1 -DSQLITE_DQS=0 -DSQLITE_LIKE_DOESNT_MATCH_BLOBS -DSQLITE_MAX_EXPR_DEPTH=0 -DSQLITE_MAX_MMAP_SIZE=0 -DSQLITE_MUTEX_APPDEF=1 -DSQLITE_MUTEX_NOOP -DSQLITE_OMIT_DECLTYPE -DSQLITE_OMIT_PROGRESS_CALLBACK -DSQLITE_OMIT_SHARED_CACHE -DSQLITE_THREADSAFE=2 -DSQLITE_USE_ALLOCA`, DO NOT EDIT.

package bin

const (
	DFTS5_TOKENIZE_AUX                     = 8
	DFTS5_TOKENIZE_DOCUMENT                = 4
	DFTS5_TOKENIZE_PREFIX                  = 2
	DFTS5_TOKENIZE_QUERY                   = 1
	DFTS5_TOKEN_COLOCATED                  = 1
	DFULLY_WITHIN                          = 2
	DHAVE_USLEEP                           = 1
	DNOT_WITHIN                            = 0
	DPARTLY_WITHIN                         = 1
	DSQLITE3_H                             = 0
	DSQLITE3_TEXT                          = 3
	DSQLITE_ABORT                          = 4
	DSQLITE_ABORT_ROLLBACK                 = 516
	DSQLITE_ACCESS_EXISTS                  = 0
	DSQLITE_ACCESS_READ                    = 2
	DSQLITE_ACCESS_READWRITE               = 1
	DSQLITE_ALTER_TABLE                    = 26
	DSQLITE_ANALYZE                        = 28
	DSQLITE_ANY                            = 5
	DSQLITE_API                            = 0
	DSQLITE_APICALL                        = 0
	DSQLITE_ATTACH                         = 24
	DSQLITE_AUTH                           = 23
	DSQLITE_AUTH_USER                      = 279
	DSQLITE_BLOB                           = 4
	DSQLITE_BUSY                           = 5
	DSQLITE_BUSY_RECOVERY                  = 261
	DSQLITE_BUSY_SNAPSHOT                  = 517
	DSQLITE_CALLBACK                       = 0
	DSQLITE_CANTOPEN                       = 14
	DSQLITE_CANTOPEN_CONVPATH              = 1038
	DSQLITE_CANTOPEN_DIRTYWAL              = 1294
	DSQLITE_CANTOPEN_FULLPATH              = 782
	DSQLITE_CANTOPEN_ISDIR                 = 526
	DSQLITE_CANTOPEN_NOTEMPDIR             = 270
	DSQLITE_CDECL                          = 0
	DSQLITE_CHECKPOINT_FULL                = 1
	DSQLITE_CHECKPOINT_PASSIVE             = 0
	DSQLITE_CHECKPOINT_RESTART             = 2
	DSQLITE_CHECKPOINT_TRUNCATE            = 3
	DSQLITE_CONFIG_COVERING_INDEX_SCAN     = 20
	DSQLITE_CONFIG_GETMALLOC               = 5
	DSQLITE_CONFIG_GETMUTEX                = 11
	DSQLITE_CONFIG_GETPCACHE               = 15
	DSQLITE_CONFIG_GETPCACHE2              = 19
	DSQLITE_CONFIG_HEAP                    = 8
	DSQLITE_CONFIG_LOG                     = 16
	DSQLITE_CONFIG_LOOKASIDE               = 13
	DSQLITE_CONFIG_MALLOC                  = 4
	DSQLITE_CONFIG_MEMDB_MAXSIZE           = 29
	DSQLITE_CONFIG_MEMSTATUS               = 9
	DSQLITE_CONFIG_MMAP_SIZE               = 22
	DSQLITE_CONFIG_MULTITHREAD             = 2
	DSQLITE_CONFIG_MUTEX                   = 10
	DSQLITE_CONFIG_PAGECACHE               = 7
	DSQLITE_CONFIG_PCACHE                  = 14
	DSQLITE_CONFIG_PCACHE2                 = 18
	DSQLITE_CONFIG_PCACHE_HDRSZ            = 24
	DSQLITE_CONFIG_PMASZ                   = 25
	DSQLITE_CONFIG_SCRATCH                 = 6
	DSQLITE_CONFIG_SERIALIZED              = 3
	DSQLITE_CONFIG_SINGLETHREAD            = 1
	DSQLITE_CONFIG_SMALL_MALLOC            = 27
	DSQLITE_CONFIG_SORTERREF_SIZE          = 28
	DSQLITE_CONFIG_SQLLOG                  = 21
	DSQLITE_CONFIG_STMTJRNL_SPILL          = 26
	DSQLITE_CONFIG_URI                     = 17
	DSQLITE_CONFIG_WIN32_HEAPSIZE          = 23
	DSQLITE_CONSTRAINT                     = 19
	DSQLITE_CONSTRAINT_CHECK               = 275
	DSQLITE_CONSTRAINT_COMMITHOOK          = 531
	DSQLITE_CONSTRAINT_FOREIGNKEY          = 787
	DSQLITE_CONSTRAINT_FUNCTION            = 1043
	DSQLITE_CONSTRAINT_NOTNULL             = 1299
	DSQLITE_CONSTRAINT_PRIMARYKEY          = 1555
	DSQLITE_CONSTRAINT_ROWID               = 2579
	DSQLITE_CONSTRAINT_TRIGGER             = 1811
	DSQLITE_CONSTRAINT_UNIQUE              = 2067
	DSQLITE_CONSTRAINT_VTAB                = 2323
	DSQLITE_COPY                           = 0
	DSQLITE_CORRUPT                        = 11
	DSQLITE_CORRUPT_SEQUENCE               = 523
	DSQLITE_CORRUPT_VTAB                   = 267
	DSQLITE_CREATE_INDEX                   = 1
	DSQLITE_CREATE_TABLE                   = 2
	DSQLITE_CREATE_TEMP_INDEX              = 3
	DSQLITE_CREATE_TEMP_TABLE              = 4
	DSQLITE_CREATE_TEMP_TRIGGER            = 5
	DSQLITE_CREATE_TEMP_VIEW               = 6
	DSQLITE_CREATE_TRIGGER                 = 7
	DSQLITE_CREATE_VIEW                    = 8
	DSQLITE_CREATE_VTABLE                  = 29
	DSQLITE_DBCONFIG_DEFENSIVE             = 1010
	DSQLITE_DBCONFIG_DQS_DDL               = 1014
	DSQLITE_DBCONFIG_DQS_DML               = 1013
	DSQLITE_DBCONFIG_ENABLE_FKEY           = 1002
	DSQLITE_DBCONFIG_ENABLE_FTS3_TOKENIZER = 1004
	DSQLITE_DBCONFIG_ENABLE_LOAD_EXTENSION = 1005
	DSQLITE_DBCONFIG_ENABLE_QPSG           = 1007
	DSQLITE_DBCONFIG_ENABLE_TRIGGER        = 1003
	DSQLITE_DBCONFIG_ENABLE_VIEW           = 1015
	DSQLITE_DBCONFIG_LEGACY_ALTER_TABLE    = 1012
	DSQLITE_DBCONFIG_LOOKASIDE             = 1001
	DSQLITE_DBCONFIG_MAINDBNAME            = 1000
	DSQLITE_DBCONFIG_MAX                   = 1015
	DSQLITE_DBCONFIG_NO_CKPT_ON_CLOSE      = 1006
	DSQLITE_DBCONFIG_RESET_DATABASE        = 1009
	DSQLITE_DBCONFIG_TRIGGER_EQP           = 1008
	DSQLITE_DBCONFIG_WRITABLE_SCHEMA       = 1011
	DSQLITE_DBSTATUS_CACHE_HIT             = 7
	DSQLITE_DBSTATUS_CACHE_MISS            = 8
	DSQLITE_DBSTATUS_CACHE_SPILL           = 12
	DSQLITE_DBSTATUS_CACHE_USED            = 1
	DSQLITE_DBSTATUS_CACHE_USED_SHARED     = 11
	DSQLITE_DBSTATUS_CACHE_WRITE           = 9
	DSQLITE_DBSTATUS_DEFERRED_FKS          = 10
	DSQLITE_DBSTATUS_LOOKASIDE_HIT         = 4
	DSQLITE_DBSTATUS_LOOKASIDE_MISS_FULL   = 6
	DSQLITE_DBSTATUS_LOOKASIDE_MISS_SIZE   = 5
	DSQLITE_DBSTATUS_LOOKASIDE_USED        = 0
	DSQLITE_DBSTATUS_MAX                   = 12
	DSQLITE_DBSTATUS_SCHEMA_USED           = 2
	DSQLITE_DBSTATUS_STMT_USED             = 3
	DSQLITE_DEFAULT_MEMSTATUS              = 0
	DSQLITE_DEFAULT_WAL_SYNCHRONOUS        = 1
	DSQLITE_DELETE                         = 9
	DSQLITE_DENY                           = 1
	DSQLITE_DEPRECATED                     = 0
	DSQLITE_DESERIALIZE_FREEONCLOSE        = 1
	DSQLITE_DESERIALIZE_READONLY           = 4
	DSQLITE_DESERIALIZE_RESIZEABLE         = 2
	DSQLITE_DETACH                         = 25
	DSQLITE_DETERMINISTIC                  = 2048
	DSQLITE_DIRECTONLY                     = 524288
	DSQLITE_DONE                           = 101
	DSQLITE_DQS                            = 0
	DSQLITE_DROP_INDEX                     = 10
	DSQLITE_DROP_TABLE                     = 11
	DSQLITE_DROP_TEMP_INDEX                = 12
	DSQLITE_DROP_TEMP_TABLE                = 13
	DSQLITE_DROP_TEMP_TRIGGER              = 14
	DSQLITE_DROP_TEMP_VIEW                 = 15
	DSQLITE_DROP_TRIGGER                   = 16
	DSQLITE_DROP_VIEW                      = 17
	DSQLITE_DROP_VTABLE                    = 30
	DSQLITE_EMPTY                          = 16
	DSQLITE_ERROR                          = 1
	DSQLITE_ERROR_MISSING_COLLSEQ          = 257
	DSQLITE_ERROR_RETRY                    = 513
	DSQLITE_ERROR_SNAPSHOT                 = 769
	DSQLITE_EXPERIMENTAL                   = 0
	DSQLITE_FAIL                           = 3
	DSQLITE_FCNTL_BEGIN_ATOMIC_WRITE       = 31
	DSQLITE_FCNTL_BUSYHANDLER              = 15
	DSQLITE_FCNTL_CHUNK_SIZE               = 6
	DSQLITE_FCNTL_COMMIT_ATOMIC_WRITE      = 32
	DSQLITE_FCNTL_COMMIT_PHASETWO          = 22
	DSQLITE_FCNTL_DATA_VERSION             = 35
	DSQLITE_FCNTL_FILE_POINTER             = 7
	DSQLITE_FCNTL_GET_LOCKPROXYFILE        = 2
	DSQLITE_FCNTL_HAS_MOVED                = 20
	DSQLITE_FCNTL_JOURNAL_POINTER          = 28
	DSQLITE_FCNTL_LAST_ERRNO               = 4
	DSQLITE_FCNTL_LOCKSTATE                = 1
	DSQLITE_FCNTL_LOCK_TIMEOUT             = 34
	DSQLITE_FCNTL_MMAP_SIZE                = 18
	DSQLITE_FCNTL_OVERWRITE                = 11
	DSQLITE_FCNTL_PDB                      = 30
	DSQLITE_FCNTL_PERSIST_WAL              = 10
	DSQLITE_FCNTL_POWERSAFE_OVERWRITE      = 13
	DSQLITE_FCNTL_PRAGMA                   = 14
	DSQLITE_FCNTL_RBU                      = 26
	DSQLITE_FCNTL_ROLLBACK_ATOMIC_WRITE    = 33
	DSQLITE_FCNTL_SET_LOCKPROXYFILE        = 3
	DSQLITE_FCNTL_SIZE_HINT                = 5
	DSQLITE_FCNTL_SIZE_LIMIT               = 36
	DSQLITE_FCNTL_SYNC                     = 21
	DSQLITE_FCNTL_SYNC_OMITTED             = 8
	DSQLITE_FCNTL_TEMPFILENAME             = 16
	DSQLITE_FCNTL_TRACE                    = 19
	DSQLITE_FCNTL_VFSNAME                  = 12
	DSQLITE_FCNTL_VFS_POINTER              = 27
	DSQLITE_FCNTL_WAL_BLOCK                = 24
	DSQLITE_FCNTL_WIN32_AV_RETRY           = 9
	DSQLITE_FCNTL_WIN32_GET_HANDLE         = 29
	DSQLITE_FCNTL_WIN32_SET_HANDLE         = 23
	DSQLITE_FCNTL_ZIPVFS                   = 25
	DSQLITE_FLOAT                          = 2
	DSQLITE_FORMAT                         = 24
	DSQLITE_FULL                           = 13
	DSQLITE_FUNCTION                       = 31
	DSQLITE_GET_LOCKPROXYFILE              = 2
	DSQLITE_IGNORE                         = 2
	DSQLITE_INDEX_CONSTRAINT_EQ            = 2
	DSQLITE_INDEX_CONSTRAINT_FUNCTION      = 150
	DSQLITE_INDEX_CONSTRAINT_GE            = 32
	DSQLITE_INDEX_CONSTRAINT_GLOB          = 66
	DSQLITE_INDEX_CONSTRAINT_GT            = 4
	DSQLITE_INDEX_CONSTRAINT_IS            = 72
	DSQLITE_INDEX_CONSTRAINT_ISNOT         = 69
	DSQLITE_INDEX_CONSTRAINT_ISNOTNULL     = 70
	DSQLITE_INDEX_CONSTRAINT_ISNULL        = 71
	DSQLITE_INDEX_CONSTRAINT_LE            = 8
	DSQLITE_INDEX_CONSTRAINT_LIKE          = 65
	DSQLITE_INDEX_CONSTRAINT_LT            = 16
	DSQLITE_INDEX_CONSTRAINT_MATCH         = 64
	DSQLITE_INDEX_CONSTRAINT_NE            = 68
	DSQLITE_INDEX_CONSTRAINT_REGEXP        = 67
	DSQLITE_INDEX_SCAN_UNIQUE              = 1
	DSQLITE_INSERT                         = 18
	DSQLITE_INTEGER                        = 1
	DSQLITE_INTERNAL                       = 2
	DSQLITE_INTERRUPT                      = 9
	DSQLITE_IOCAP_ATOMIC                   = 1
	DSQLITE_IOCAP_ATOMIC16K                = 64
	DSQLITE_IOCAP_ATOMIC1K                 = 4
	DSQLITE_IOCAP_ATOMIC2K                 = 8
	DSQLITE_IOCAP_ATOMIC32K                = 128
	DSQLITE_IOCAP_ATOMIC4K                 = 16
	DSQLITE_IOCAP_ATOMIC512                = 2
	DSQLITE_IOCAP_ATOMIC64K                = 256
	DSQLITE_IOCAP_ATOMIC8K                 = 32
	DSQLITE_IOCAP_BATCH_ATOMIC             = 16384
	DSQLITE_IOCAP_IMMUTABLE                = 8192
	DSQLITE_IOCAP_POWERSAFE_OVERWRITE      = 4096
	DSQLITE_IOCAP_SAFE_APPEND              = 512
	DSQLITE_IOCAP_SEQUENTIAL               = 1024
	DSQLITE_IOCAP_UNDELETABLE_WHEN_OPEN    = 2048
	DSQLITE_IOERR                          = 10
	DSQLITE_IOERR_ACCESS                   = 3338
	DSQLITE_IOERR_AUTH                     = 7178
	DSQLITE_IOERR_BEGIN_ATOMIC             = 7434
	DSQLITE_IOERR_BLOCKED                  = 2826
	DSQLITE_IOERR_CHECKRESERVEDLOCK        = 3594
	DSQLITE_IOERR_CLOSE                    = 4106
	DSQLITE_IOERR_COMMIT_ATOMIC            = 7690
	DSQLITE_IOERR_CONVPATH                 = 6666
	DSQLITE_IOERR_DELETE                   = 2570
	DSQLITE_IOERR_DELETE_NOENT             = 5898
	DSQLITE_IOERR_DIR_CLOSE                = 4362
	DSQLITE_IOERR_DIR_FSYNC                = 1290
	DSQLITE_IOERR_FSTAT                    = 1802
	DSQLITE_IOERR_FSYNC                    = 1034
	DSQLITE_IOERR_GETTEMPPATH              = 6410
	DSQLITE_IOERR_LOCK                     = 3850
	DSQLITE_IOERR_MMAP                     = 6154
	DSQLITE_IOERR_NOMEM                    = 3082
	DSQLITE_IOERR_RDLOCK                   = 2314
	DSQLITE_IOERR_READ                     = 266
	DSQLITE_IOERR_ROLLBACK_ATOMIC          = 7946
	DSQLITE_IOERR_SEEK                     = 5642
	DSQLITE_IOERR_SHMLOCK                  = 5130
	DSQLITE_IOERR_SHMMAP                   = 5386
	DSQLITE_IOERR_SHMOPEN                  = 4618
	DSQLITE_IOERR_SHMSIZE                  = 4874
	DSQLITE_IOERR_SHORT_READ               = 522
	DSQLITE_IOERR_TRUNCATE                 = 1546
	DSQLITE_IOERR_UNLOCK                   = 2058
	DSQLITE_IOERR_VNODE                    = 6922
	DSQLITE_IOERR_WRITE                    = 778
	DSQLITE_LAST_ERRNO                     = 4
	DSQLITE_LIKE_DOESNT_MATCH_BLOBS        = 1
	DSQLITE_LIMIT_ATTACHED                 = 7
	DSQLITE_LIMIT_COLUMN                   = 2
	DSQLITE_LIMIT_COMPOUND_SELECT          = 4
	DSQLITE_LIMIT_EXPR_DEPTH               = 3
	DSQLITE_LIMIT_FUNCTION_ARG             = 6
	DSQLITE_LIMIT_LENGTH                   = 0
	DSQLITE_LIMIT_LIKE_PATTERN_LENGTH      = 8
	DSQLITE_LIMIT_SQL_LENGTH               = 1
	DSQLITE_LIMIT_TRIGGER_DEPTH            = 10
	DSQLITE_LIMIT_VARIABLE_NUMBER          = 9
	DSQLITE_LIMIT_VDBE_OP                  = 5
	DSQLITE_LIMIT_WORKER_THREADS           = 11
	DSQLITE_LOCKED                         = 6
	DSQLITE_LOCKED_SHAREDCACHE             = 262
	DSQLITE_LOCKED_VTAB                    = 518
	DSQLITE_LOCK_EXCLUSIVE                 = 4
	DSQLITE_LOCK_NONE                      = 0
	DSQLITE_LOCK_PENDING                   = 3
	DSQLITE_LOCK_RESERVED                  = 2
	DSQLITE_LOCK_SHARED                    = 1
	DSQLITE_MAX_EXPR_DEPTH                 = 0
	DSQLITE_MAX_MMAP_SIZE                  = 0
	DSQLITE_MISMATCH                       = 20
	DSQLITE_MISUSE                         = 21
	DSQLITE_MUTEX_APPDEF                   = 1
	DSQLITE_MUTEX_FAST                     = 0
	DSQLITE_MUTEX_NOOP                     = 1
	DSQLITE_MUTEX_RECURSIVE                = 1
	DSQLITE_MUTEX_STATIC_APP1              = 8
	DSQLITE_MUTEX_STATIC_APP2              = 9
	DSQLITE_MUTEX_STATIC_APP3              = 10
	DSQLITE_MUTEX_STATIC_LRU               = 6
	DSQLITE_MUTEX_STATIC_LRU2              = 7
	DSQLITE_MUTEX_STATIC_MASTER            = 2
	DSQLITE_MUTEX_STATIC_MEM               = 3
	DSQLITE_MUTEX_STATIC_MEM2              = 4
	DSQLITE_MUTEX_STATIC_OPEN              = 4
	DSQLITE_MUTEX_STATIC_PMEM              = 7
	DSQLITE_MUTEX_STATIC_PRNG              = 5
	DSQLITE_MUTEX_STATIC_VFS1              = 11
	DSQLITE_MUTEX_STATIC_VFS2              = 12
	DSQLITE_MUTEX_STATIC_VFS3              = 13
	DSQLITE_NOLFS                          = 22
	DSQLITE_NOMEM                          = 7
	DSQLITE_NOTADB                         = 26
	DSQLITE_NOTFOUND                       = 12
	DSQLITE_NOTICE                         = 27
	DSQLITE_NOTICE_RECOVER_ROLLBACK        = 539
	DSQLITE_NOTICE_RECOVER_WAL             = 283
	DSQLITE_NULL                           = 5
	DSQLITE_OK                             = 0
	DSQLITE_OK_LOAD_PERMANENTLY            = 256
	DSQLITE_OMIT_DECLTYPE                  = 1
	DSQLITE_OMIT_PROGRESS_CALLBACK         = 1
	DSQLITE_OMIT_SHARED_CACHE              = 1
	DSQLITE_OPEN_AUTOPROXY                 = 32
	DSQLITE_OPEN_CREATE                    = 4
	DSQLITE_OPEN_DELETEONCLOSE             = 8
	DSQLITE_OPEN_EXCLUSIVE                 = 16
	DSQLITE_OPEN_FULLMUTEX                 = 65536
	DSQLITE_OPEN_MAIN_DB                   = 256
	DSQLITE_OPEN_MAIN_JOURNAL              = 2048
	DSQLITE_OPEN_MASTER_JOURNAL            = 16384
	DSQLITE_OPEN_MEMORY                    = 128
	DSQLITE_OPEN_NOMUTEX                   = 32768
	DSQLITE_OPEN_PRIVATECACHE              = 262144
	DSQLITE_OPEN_READONLY                  = 1
	DSQLITE_OPEN_READWRITE                 = 2
	DSQLITE_OPEN_SHAREDCACHE               = 131072
	DSQLITE_OPEN_SUBJOURNAL                = 8192
	DSQLITE_OPEN_TEMP_DB                   = 512
	DSQLITE_OPEN_TEMP_JOURNAL              = 4096
	DSQLITE_OPEN_TRANSIENT_DB              = 1024
	DSQLITE_OPEN_URI                       = 64
	DSQLITE_OPEN_WAL                       = 524288
	DSQLITE_PERM                           = 3
	DSQLITE_PRAGMA                         = 19
	DSQLITE_PREPARE_NORMALIZE              = 2
	DSQLITE_PREPARE_NO_VTAB                = 4
	DSQLITE_PREPARE_PERSISTENT             = 1
	DSQLITE_PROTOCOL                       = 15
	DSQLITE_RANGE                          = 25
	DSQLITE_READ                           = 20
	DSQLITE_READONLY                       = 8
	DSQLITE_READONLY_CANTINIT              = 1288
	DSQLITE_READONLY_CANTLOCK              = 520
	DSQLITE_READONLY_DBMOVED               = 1032
	DSQLITE_READONLY_DIRECTORY             = 1544
	DSQLITE_READONLY_RECOVERY              = 264
	DSQLITE_READONLY_ROLLBACK              = 776
	DSQLITE_RECURSIVE                      = 33
	DSQLITE_REINDEX                        = 27
	DSQLITE_REPLACE                        = 5
	DSQLITE_ROLLBACK                       = 1
	DSQLITE_ROW                            = 100
	DSQLITE_SAVEPOINT                      = 32
	DSQLITE_SCANSTAT_EST                   = 2
	DSQLITE_SCANSTAT_EXPLAIN               = 4
	DSQLITE_SCANSTAT_NAME                  = 3
	DSQLITE_SCANSTAT_NLOOP                 = 0
	DSQLITE_SCANSTAT_NVISIT                = 1
	DSQLITE_SCANSTAT_SELECTID              = 5
	DSQLITE_SCHEMA                         = 17
	DSQLITE_SELECT                         = 21
	DSQLITE_SERIALIZE_NOCOPY               = 1
	DSQLITE_SET_LOCKPROXYFILE              = 3
	DSQLITE_SHM_EXCLUSIVE                  = 8
	DSQLITE_SHM_LOCK                       = 2
	DSQLITE_SHM_NLOCK                      = 8
	DSQLITE_SHM_SHARED                     = 4
	DSQLITE_SHM_UNLOCK                     = 1
	DSQLITE_STATUS_MALLOC_COUNT            = 9
	DSQLITE_STATUS_MALLOC_SIZE             = 5
	DSQLITE_STATUS_MEMORY_USED             = 0
	DSQLITE_STATUS_PAGECACHE_OVERFLOW      = 2
	DSQLITE_STATUS_PAGECACHE_SIZE          = 7
	DSQLITE_STATUS_PAGECACHE_USED          = 1
	DSQLITE_STATUS_PARSER_STACK            = 6
	DSQLITE_STATUS_SCRATCH_OVERFLOW        = 4
	DSQLITE_STATUS_SCRATCH_SIZE            = 8
	DSQLITE_STATUS_SCRATCH_USED            = 3
	DSQLITE_STDCALL                        = 0
	DSQLITE_STMTSTATUS_AUTOINDEX           = 3
	DSQLITE_STMTSTATUS_FULLSCAN_STEP       = 1
	DSQLITE_STMTSTATUS_MEMUSED             = 99
	DSQLITE_STMTSTATUS_REPREPARE           = 5
	DSQLITE_STMTSTATUS_RUN                 = 6
	DSQLITE_STMTSTATUS_SORT                = 2
	DSQLITE_STMTSTATUS_VM_STEP             = 4
	DSQLITE_SUBTYPE                        = 1048576
	DSQLITE_SYNC_DATAONLY                  = 16
	DSQLITE_SYNC_FULL                      = 3
	DSQLITE_SYNC_NORMAL                    = 2
	DSQLITE_SYSAPI                         = 0
	DSQLITE_TESTCTRL_ALWAYS                = 13
	DSQLITE_TESTCTRL_ASSERT                = 12
	DSQLITE_TESTCTRL_BENIGN_MALLOC_HOOKS   = 10
	DSQLITE_TESTCTRL_BITVEC_TEST           = 8
	DSQLITE_TESTCTRL_BYTEORDER             = 22
	DSQLITE_TESTCTRL_EXPLAIN_STMT          = 19
	DSQLITE_TESTCTRL_EXTRA_SCHEMA_CHECKS   = 29
	DSQLITE_TESTCTRL_FAULT_INSTALL         = 9
	DSQLITE_TESTCTRL_FIRST                 = 5
	DSQLITE_TESTCTRL_IMPOSTER              = 25
	DSQLITE_TESTCTRL_INTERNAL_FUNCTIONS    = 17
	DSQLITE_TESTCTRL_ISINIT                = 23
	DSQLITE_TESTCTRL_ISKEYWORD             = 16
	DSQLITE_TESTCTRL_LAST                  = 29
	DSQLITE_TESTCTRL_LOCALTIME_FAULT       = 18
	DSQLITE_TESTCTRL_NEVER_CORRUPT         = 20
	DSQLITE_TESTCTRL_ONCE_RESET_THRESHOLD  = 19
	DSQLITE_TESTCTRL_OPTIMIZATIONS         = 15
	DSQLITE_TESTCTRL_PARSER_COVERAGE       = 26
	DSQLITE_TESTCTRL_PENDING_BYTE          = 11
	DSQLITE_TESTCTRL_PRNG_RESET            = 7
	DSQLITE_TESTCTRL_PRNG_RESTORE          = 6
	DSQLITE_TESTCTRL_PRNG_SAVE             = 5
	DSQLITE_TESTCTRL_PRNG_SEED             = 28
	DSQLITE_TESTCTRL_RESERVE               = 14
	DSQLITE_TESTCTRL_RESULT_INTREAL        = 27
	DSQLITE_TESTCTRL_SCRATCHMALLOC         = 17
	DSQLITE_TESTCTRL_SORTER_MMAP           = 24
	DSQLITE_TESTCTRL_VDBE_COVERAGE         = 21
	DSQLITE_TEXT                           = 3
	DSQLITE_THREADSAFE                     = 2
	DSQLITE_TOOBIG                         = 18
	DSQLITE_TRACE_CLOSE                    = 8
	DSQLITE_TRACE_PROFILE                  = 2
	DSQLITE_TRACE_ROW                      = 4
	DSQLITE_TRACE_STMT                     = 1
	DSQLITE_TRANSACTION                    = 22
	DSQLITE_UPDATE                         = 23
	DSQLITE_USE_ALLOCA                     = 1
	DSQLITE_UTF16                          = 4
	DSQLITE_UTF16BE                        = 3
	DSQLITE_UTF16LE                        = 2
	DSQLITE_UTF16_ALIGNED                  = 8
	DSQLITE_UTF8                           = 1
	DSQLITE_VERSION_NUMBER                 = 3030001
	DSQLITE_VTAB_CONSTRAINT_SUPPORT        = 1
	DSQLITE_WARNING                        = 28
	DSQLITE_WARNING_AUTOINDEX              = 284
	DSQLITE_WIN32_DATA_DIRECTORY_TYPE      = 1
	DSQLITE_WIN32_TEMP_DIRECTORY_TYPE      = 2
	D_ANSI_STDARG_H_                       = 0
	D_FTS5_H                               = 0
	D_LP64                                 = 1
	D_SQLITE3RTREE_H_                      = 0
	D_STDARG_H                             = 0
	D_STDC_PREDEF_H                        = 1
	D_VA_LIST                              = 0
	D_VA_LIST_                             = 0
	D_VA_LIST_DEFINED                      = 0
	D_VA_LIST_T_H                          = 0
	D__ATOMIC_ACQUIRE                      = 2
	D__ATOMIC_ACQ_REL                      = 4
	D__ATOMIC_CONSUME                      = 1
	D__ATOMIC_HLE_ACQUIRE                  = 65536
	D__ATOMIC_HLE_RELEASE                  = 131072
	D__ATOMIC_RELAXED                      = 0
	D__ATOMIC_RELEASE                      = 3
	D__ATOMIC_SEQ_CST                      = 5
	D__BIGGEST_ALIGNMENT__                 = 16
	D__BYTE_ORDER__                        = 1234
	D__CHAR_BIT__                          = 8
	D__COUNTER__                           = 0
	D__DBL_DECIMAL_DIG__                   = 17
	D__DBL_DIG__                           = 15
	D__DBL_HAS_DENORM__                    = 1
	D__DBL_HAS_INFINITY__                  = 1
	D__DBL_HAS_QUIET_NAN__                 = 1
	D__DBL_MANT_DIG__                      = 53
	D__DBL_MAX_10_EXP__                    = 308
	D__DBL_MAX_EXP__                       = 1024
	D__DBL_MIN_10_EXP__                    = -307
	D__DBL_MIN_EXP__                       = -1021
	D__DEC128_MANT_DIG__                   = 34
	D__DEC128_MAX_EXP__                    = 6145
	D__DEC128_MIN_EXP__                    = -6142
	D__DEC32_MANT_DIG__                    = 7
	D__DEC32_MAX_EXP__                     = 97
	D__DEC32_MIN_EXP__                     = -94
	D__DEC64_MANT_DIG__                    = 16
	D__DEC64_MAX_EXP__                     = 385
	D__DEC64_MIN_EXP__                     = -382
	D__DECIMAL_BID_FORMAT__                = 1
	D__DECIMAL_DIG__                       = 21
	D__DEC_EVAL_METHOD__                   = 2
	D__DI__                                = 0
	D__ELF__                               = 1
	D__FINITE_MATH_ONLY__                  = 0
	D__FLOAT_WORD_ORDER__                  = 1234
	D__FLT128_DECIMAL_DIG__                = 36
	D__FLT128_DIG__                        = 33
	D__FLT128_HAS_DENORM__                 = 1
	D__FLT128_HAS_INFINITY__               = 1
	D__FLT128_HAS_QUIET_NAN__              = 1
	D__FLT128_MANT_DIG__                   = 113
	D__FLT128_MAX_10_EXP__                 = 4932
	D__FLT128_MAX_EXP__                    = 16384
	D__FLT128_MIN_10_EXP__                 = -4931
	D__FLT128_MIN_EXP__                    = -16381
	D__FLT32X_DECIMAL_DIG__                = 17
	D__FLT32X_DIG__                        = 15
	D__FLT32X_HAS_DENORM__                 = 1
	D__FLT32X_HAS_INFINITY__               = 1
	D__FLT32X_HAS_QUIET_NAN__              = 1
	D__FLT32X_MANT_DIG__                   = 53
	D__FLT32X_MAX_10_EXP__                 = 308
	D__FLT32X_MAX_EXP__                    = 1024
	D__FLT32X_MIN_10_EXP__                 = -307
	D__FLT32X_MIN_EXP__                    = -1021
	D__FLT32_DECIMAL_DIG__                 = 9
	D__FLT32_DIG__                         = 6
	D__FLT32_HAS_DENORM__                  = 1
	D__FLT32_HAS_INFINITY__                = 1
	D__FLT32_HAS_QUIET_NAN__               = 1
	D__FLT32_MANT_DIG__                    = 24
	D__FLT32_MAX_10_EXP__                  = 38
	D__FLT32_MAX_EXP__                     = 128
	D__FLT32_MIN_10_EXP__                  = -37
	D__FLT32_MIN_EXP__                     = -125
	D__FLT64X_DECIMAL_DIG__                = 21
	D__FLT64X_DIG__                        = 18
	D__FLT64X_HAS_DENORM__                 = 1
	D__FLT64X_HAS_INFINITY__               = 1
	D__FLT64X_HAS_QUIET_NAN__              = 1
	D__FLT64X_MANT_DIG__                   = 64
	D__FLT64X_MAX_10_EXP__                 = 4932
	D__FLT64X_MAX_EXP__                    = 16384
	D__FLT64X_MIN_10_EXP__                 = -4931
	D__FLT64X_MIN_EXP__                    = -16381
	D__FLT64_DECIMAL_DIG__                 = 17
	D__FLT64_DIG__                         = 15
	D__FLT64_HAS_DENORM__                  = 1
	D__FLT64_HAS_INFINITY__                = 1
	D__FLT64_HAS_QUIET_NAN__               = 1
	D__FLT64_MANT_DIG__                    = 53
	D__FLT64_MAX_10_EXP__                  = 308
	D__FLT64_MAX_EXP__                     = 1024
	D__FLT64_MIN_10_EXP__                  = -307
	D__FLT64_MIN_EXP__                     = -1021
	D__FLT_DECIMAL_DIG__                   = 9
	D__FLT_DIG__                           = 6
	D__FLT_EVAL_METHOD_TS_18661_3__        = 0
	D__FLT_EVAL_METHOD__                   = 0
	D__FLT_HAS_DENORM__                    = 1
	D__FLT_HAS_INFINITY__                  = 1
	D__FLT_HAS_QUIET_NAN__                 = 1
	D__FLT_MANT_DIG__                      = 24
	D__FLT_MAX_10_EXP__                    = 38
	D__FLT_MAX_EXP__                       = 128
	D__FLT_MIN_10_EXP__                    = -37
	D__FLT_MIN_EXP__                       = -125
	D__FLT_RADIX__                         = 2
	D__FXSR__                              = 1
	D__GNUC_VA_LIST                        = 0
	D__GXX_ABI_VERSION                     = 1011
	D__HI__                                = 0
	D__INT16_MAX__                         = 32767
	D__INT32_MAX__                         = 2147483647
	D__INT64_MAX__                         = 9223372036854775807
	D__INT8_MAX__                          = 127
	D__INTMAX_MAX__                        = 9223372036854775807
	D__INTMAX_WIDTH__                      = 64
	D__INTPTR_MAX__                        = 9223372036854775807
	D__INTPTR_WIDTH__                      = 64
	D__INT_FAST16_MAX__                    = 9223372036854775807
	D__INT_FAST16_WIDTH__                  = 64
	D__INT_FAST32_MAX__                    = 9223372036854775807
	D__INT_FAST32_WIDTH__                  = 64
	D__INT_FAST64_MAX__                    = 9223372036854775807
	D__INT_FAST64_WIDTH__                  = 64
	D__INT_FAST8_MAX__                     = 127
	D__INT_FAST8_WIDTH__                   = 8
	D__INT_LEAST16_MAX__                   = 32767
	D__INT_LEAST16_WIDTH__                 = 16
	D__INT_LEAST32_MAX__                   = 2147483647
	D__INT_LEAST32_WIDTH__                 = 32
	D__INT_LEAST64_MAX__                   = 9223372036854775807
	D__INT_LEAST64_WIDTH__                 = 64
	D__INT_LEAST8_MAX__                    = 127
	D__INT_LEAST8_WIDTH__                  = 8
	D__INT_MAX__                           = 2147483647
	D__INT_WIDTH__                         = 32
	D__LDBL_DECIMAL_DIG__                  = 21
	D__LDBL_DIG__                          = 18
	D__LDBL_HAS_DENORM__                   = 1
	D__LDBL_HAS_INFINITY__                 = 1
	D__LDBL_HAS_QUIET_NAN__                = 1
	D__LDBL_MANT_DIG__                     = 64
	D__LDBL_MAX_10_EXP__                   = 4932
	D__LDBL_MAX_EXP__                      = 16384
	D__LDBL_MIN_10_EXP__                   = -4931
	D__LDBL_MIN_EXP__                      = -16381
	D__LINE__                              = 0
	D__LONG_LONG_MAX__                     = 9223372036854775807
	D__LONG_LONG_WIDTH__                   = 64
	D__LONG_MAX__                          = 9223372036854775807
	D__LONG_WIDTH__                        = 64
	D__LP64__                              = 1
	D__MMX__                               = 1
	D__NO_INLINE__                         = 1
	D__ORDER_BIG_ENDIAN__                  = 4321
	D__ORDER_LITTLE_ENDIAN__               = 1234
	D__ORDER_PDP_ENDIAN__                  = 3412
	D__PRAGMA_REDEFINE_EXTNAME             = 1
	D__PTRDIFF_MAX__                       = 9223372036854775807
	D__PTRDIFF_WIDTH__                     = 64
	D__QI__                                = 0
	D__REGISTER_PREFIX__                   = 0
	D__SCHAR_MAX__                         = 127
	D__SCHAR_WIDTH__                       = 8
	D__SEG_FS                              = 1
	D__SEG_GS                              = 1
	D__SHRT_MAX__                          = 32767
	D__SHRT_WIDTH__                        = 16
	D__SIG_ATOMIC_MAX__                    = 2147483647
	D__SIG_ATOMIC_MIN__                    = -2147483648
	D__SIG_ATOMIC_WIDTH__                  = 32
	D__SIZEOF_DOUBLE__                     = 8
	D__SIZEOF_FLOAT128__                   = 16
	D__SIZEOF_FLOAT80__                    = 16
	D__SIZEOF_FLOAT__                      = 4
	D__SIZEOF_INT128__                     = 16
	D__SIZEOF_INT__                        = 4
	D__SIZEOF_LONG_DOUBLE__                = 16
	D__SIZEOF_LONG_LONG__                  = 8
	D__SIZEOF_LONG__                       = 8
	D__SIZEOF_POINTER__                    = 8
	D__SIZEOF_PTRDIFF_T__                  = 8
	D__SIZEOF_SHORT__                      = 2
	D__SIZEOF_SIZE_T__                     = 8
	D__SIZEOF_WCHAR_T__                    = 4
	D__SIZEOF_WINT_T__                     = 4
	D__SIZE_MAX__                          = 18446744073709551615
	D__SIZE_WIDTH__                        = 64
	D__SI__                                = 0
	D__SSE2_MATH__                         = 1
	D__SSE2__                              = 1
	D__SSE_MATH__                          = 1
	D__SSE__                               = 1
	D__STDC_HOSTED__                       = 1
	D__STDC_IEC_559_COMPLEX__              = 1
	D__STDC_IEC_559__                      = 1
	D__STDC_ISO_10646__                    = 201706
	D__STDC_NO_THREADS__                   = 1
	D__STDC_UTF_16__                       = 1
	D__STDC_UTF_32__                       = 1
	D__STDC_VERSION__                      = 201112
	D__STDC__                              = 1
	D__UINT16_MAX__                        = 65535
	D__UINT32_MAX__                        = 4294967295
	D__UINT64_MAX__                        = 18446744073709551615
	D__UINT8_MAX__                         = 255
	D__UINTMAX_MAX__                       = 18446744073709551615
	D__UINTPTR_MAX__                       = 18446744073709551615
	D__UINT_FAST16_MAX__                   = 18446744073709551615
	D__UINT_FAST32_MAX__                   = 18446744073709551615
	D__UINT_FAST64_MAX__                   = 18446744073709551615
	D__UINT_FAST8_MAX__                    = 255
	D__UINT_LEAST16_MAX__                  = 65535
	D__UINT_LEAST32_MAX__                  = 4294967295
	D__UINT_LEAST64_MAX__                  = 18446744073709551615
	D__UINT_LEAST8_MAX__                   = 255
	D__USER_LABEL_PREFIX__                 = 0
	D__WCHAR_MAX__                         = 2147483647
	D__WCHAR_MIN__                         = -2147483648
	D__WCHAR_WIDTH__                       = 32
	D__WINT_MAX__                          = 4294967295
	D__WINT_MIN__                          = 0
	D__WINT_WIDTH__                        = 32
	D__amd64                               = 1
	D__amd64__                             = 1
	D__code_model_small__                  = 1
	D__extension__                         = 0
	D__k8                                  = 1
	D__k8__                                = 1
	D__linux                               = 1
	D__linux__                             = 1
	D__unix                                = 1
	D__unix__                              = 1
	D__va_list__                           = 0
	D__word__                              = 0
	D__x86_64                              = 1
	D__x86_64__                            = 1
	Dlinux                                 = 1
	Dunix                                  = 1
)

type SFts5ExtensionApi = struct {
	FiVersion           int32
	FxUserData          uintptr
	FxColumnCount       uintptr
	FxRowCount          uintptr
	FxColumnTotalSize   uintptr
	FxTokenize          uintptr
	FxPhraseCount       uintptr
	FxPhraseSize        uintptr
	FxInstCount         uintptr
	FxInst              uintptr
	FxRowid             uintptr
	FxColumnText        uintptr
	FxColumnSize        uintptr
	FxQueryPhrase       uintptr
	FxSetAuxdata        uintptr
	FxGetAuxdata        uintptr
	FxPhraseFirst       uintptr
	FxPhraseNext        uintptr
	FxPhraseFirstColumn uintptr
	FxPhraseNextColumn  uintptr
}

type SFts5PhraseIter = struct {
	Fa uintptr
	Fb uintptr
}

type Sfts5_api = struct {
	FiVersion         int32
	FxCreateTokenizer uintptr
	FxFindTokenizer   uintptr
	FxCreateFunction  uintptr
}

type Sfts5_tokenizer = struct {
	FxCreate   uintptr
	FxDelete   uintptr
	FxTokenize uintptr
}

type Ssqlite3_file = struct{ FpMethods uintptr }

type Ssqlite3_io_methods = struct {
	FiVersion               int32
	FxClose                 uintptr
	FxRead                  uintptr
	FxWrite                 uintptr
	FxTruncate              uintptr
	FxSync                  uintptr
	FxFileSize              uintptr
	FxLock                  uintptr
	FxUnlock                uintptr
	FxCheckReservedLock     uintptr
	FxFileControl           uintptr
	FxSectorSize            uintptr
	FxDeviceCharacteristics uintptr
	FxShmMap                uintptr
	FxShmLock               uintptr
	FxShmBarrier            uintptr
	FxShmUnmap              uintptr
	FxFetch                 uintptr
	FxUnfetch               uintptr
}

type Ssqlite3_mem_methods = struct {
	FxMalloc   uintptr
	FxFree     uintptr
	FxRealloc  uintptr
	FxSize     uintptr
	FxRoundup  uintptr
	FxInit     uintptr
	FxShutdown uintptr
	FpAppData  uintptr
}

type Ssqlite3_module = struct {
	FiVersion      int32
	FxCreate       uintptr
	FxConnect      uintptr
	FxBestIndex    uintptr
	FxDisconnect   uintptr
	FxDestroy      uintptr
	FxOpen         uintptr
	FxClose        uintptr
	FxFilter       uintptr
	FxNext         uintptr
	FxEof          uintptr
	FxColumn       uintptr
	FxRowid        uintptr
	FxUpdate       uintptr
	FxBegin        uintptr
	FxSync         uintptr
	FxCommit       uintptr
	FxRollback     uintptr
	FxFindFunction uintptr
	FxRename       uintptr
	FxSavepoint    uintptr
	FxRelease      uintptr
	FxRollbackTo   uintptr
	FxShadowName   uintptr
}

type Ssqlite3_mutex_methods = struct {
	FxMutexInit    uintptr
	FxMutexEnd     uintptr
	FxMutexAlloc   uintptr
	FxMutexFree    uintptr
	FxMutexEnter   uintptr
	FxMutexTry     uintptr
	FxMutexLeave   uintptr
	FxMutexHeld    uintptr
	FxMutexNotheld uintptr
}

type Ssqlite3_pcache_methods = struct {
	FpArg       uintptr
	FxInit      uintptr
	FxShutdown  uintptr
	FxCreate    uintptr
	FxCachesize uintptr
	FxPagecount uintptr
	FxFetch     uintptr
	FxUnpin     uintptr
	FxRekey     uintptr
	FxTruncate  uintptr
	FxDestroy   uintptr
}

type Ssqlite3_pcache_methods2 = struct {
	FiVersion   int32
	FpArg       uintptr
	FxInit      uintptr
	FxShutdown  uintptr
	FxCreate    uintptr
	FxCachesize uintptr
	FxPagecount uintptr
	FxFetch     uintptr
	FxUnpin     uintptr
	FxRekey     uintptr
	FxTruncate  uintptr
	FxDestroy   uintptr
	FxShrink    uintptr
}

type Ssqlite3_pcache_page = struct {
	FpBuf   uintptr
	FpExtra uintptr
}

type Ssqlite3_rtree_geometry = struct {
	FpContext uintptr
	FnParam   int32
	FaParam   uintptr
	FpUser    uintptr
	FxDelUser uintptr
}

type Ssqlite3_snapshot = struct{ Fhidden [48]int8 }

type Ssqlite3_vfs = struct {
	FiVersion          int32
	FszOsFile          int32
	FmxPathname        int32
	FpNext             uintptr
	FzName             uintptr
	FpAppData          uintptr
	FxOpen             uintptr
	FxDelete           uintptr
	FxAccess           uintptr
	FxFullPathname     uintptr
	FxDlOpen           uintptr
	FxDlError          uintptr
	FxDlSym            uintptr
	FxDlClose          uintptr
	FxRandomness       uintptr
	FxSleep            uintptr
	FxCurrentTime      uintptr
	FxGetLastError     uintptr
	FxCurrentTimeInt64 uintptr
	FxSetSystemCall    uintptr
	FxGetSystemCall    uintptr
	FxNextSystemCall   uintptr
}

type Ssqlite3_vtab = struct {
	FpModule uintptr
	FnRef    int32
	FzErrMsg uintptr
}

type Ssqlite3_vtab_cursor = struct{ FpVtab uintptr }

var sbin__ [1]byte
