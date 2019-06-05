// DO NOT EDIT.
// GENERATED by go:generate at 2019-06-05 15:12:15.983028 +0000 UTC.
package gounity

type BlockHostAccessEnum int

const (

	// BlockHostAccessOff means: Access is disabled.
	BlockHostAccessOff BlockHostAccessEnum = 0

	// BlockHostAccessOn means: Access is enabled.
	BlockHostAccessOn BlockHostAccessEnum = 1

	// BlockHostAccessMixed means: (Applies to Consistency Groups only.) Indicates that LUNs in a consistency group have different host access.
	BlockHostAccessMixed BlockHostAccessEnum = 2
)

type FSSupportedProtocolEnum int

const (

	// FSSupportedProtocolNFS means: Only NFS share protocol supported by file system.
	FSSupportedProtocolNFS FSSupportedProtocolEnum = 0

	// FSSupportedProtocolCIFS means: Only SMB (also known as CIFS) share protocol is supported by file system.
	FSSupportedProtocolCIFS FSSupportedProtocolEnum = 1

	// FSSupportedProtocolMultiprotocol means: Both share protocols NFS and SMB (also known as CIFS) are supported by file system.
	FSSupportedProtocolMultiprotocol FSSupportedProtocolEnum = 2
)

type HostLUNAccessEnum int

const (

	// HostLUNAccessNoAccess means: No access.
	HostLUNAccessNoAccess HostLUNAccessEnum = 0

	// HostLUNAccessProduction means: Access to production LUNs only.
	HostLUNAccessProduction HostLUNAccessEnum = 1

	// HostLUNAccessSnapshot means: Access to LUN snapshots only.
	HostLUNAccessSnapshot HostLUNAccessEnum = 2

	// HostLUNAccessBoth means: Access to both production LUNs and their snapshots.
	HostLUNAccessBoth HostLUNAccessEnum = 3

	// HostLUNAccessProductionOn means: Request to grant production access to LUNs for host. Should be used only by GUI.
	HostLUNAccessProductionOn HostLUNAccessEnum = 4

	// HostLUNAccessProductionOff means: Request to deny production access to LUNs for host. Should be used only by GUI.
	HostLUNAccessProductionOff HostLUNAccessEnum = 5

	// HostLUNAccessMixed means: (Applies to consistency groups only.) Indicates that LUNs in a consistency group have different host access. Do not use this value in Create or Modify requests.
	HostLUNAccessMixed HostLUNAccessEnum = 65535
)

type HostLunTypeEnum int

const (

	// HostLunTypeUnknown means: Unknown LUN type.
	HostLunTypeUnknown HostLunTypeEnum = 0

	// HostLunTypeLUN means: Production LUN.
	HostLunTypeLUN HostLunTypeEnum = 1

	// HostLunTypeLUN_Snap means: Snapshot LUN.
	HostLunTypeLUN_Snap HostLunTypeEnum = 2
)

type IpProtocolVersionEnum int

const (

	// IpProtocolVersionIPv4 means: Network interface uses IPv4 address.
	IpProtocolVersionIPv4 IpProtocolVersionEnum = 4

	// IpProtocolVersionIPv6 means: Network interface uses IPv6 address.
	IpProtocolVersionIPv6 IpProtocolVersionEnum = 6
)

type JobStateEnum int

const (

	// JobStateQueued means: Job is queued to run.
	JobStateQueued JobStateEnum = 1

	// JobStateRunning means: Job is running.
	JobStateRunning JobStateEnum = 2

	// JobStateSuspended means: Job is suspended.
	JobStateSuspended JobStateEnum = 3

	// JobStateCompleted means: Job completed successfully.
	JobStateCompleted JobStateEnum = 4

	// JobStateFailed means: Job is failed, interrupted or terminated.
	JobStateFailed JobStateEnum = 5

	// JobStateRolling_Back means: Job has failed and is attempting to roll back.
	JobStateRolling_Back JobStateEnum = 6

	// JobStateCompleted_With_Problems means: Job ran to the end, but a task returned an error.
	JobStateCompleted_With_Problems JobStateEnum = 7
)

type JobTaskStateEnum int

const (

	// JobTaskStateNot_Started means: Job task is waiting to run.
	JobTaskStateNot_Started JobTaskStateEnum = 0

	// JobTaskStateRunning means: Job task is running.
	JobTaskStateRunning JobTaskStateEnum = 1

	// JobTaskStateCompleted means: Job task completed successfully.
	JobTaskStateCompleted JobTaskStateEnum = 2

	// JobTaskStateFailed means: Job task failed.
	JobTaskStateFailed JobTaskStateEnum = 3

	// JobTaskStateRolling_Back means: Job task is rolling back.
	JobTaskStateRolling_Back JobTaskStateEnum = 4

	// JobTaskStateCompleted_With_Problems means: Job ran to the end, but a task returned an error.
	JobTaskStateCompleted_With_Problems JobTaskStateEnum = 5

	// JobTaskStateSuspended means: Job is suspended.
	JobTaskStateSuspended JobTaskStateEnum = 6
)

type NFSShareDefaultAccessEnum int

const (

	// NFSShareDefaultAccessNoAccess means: Deny access to the share for the hosts.
	NFSShareDefaultAccessNoAccess NFSShareDefaultAccessEnum = 0

	// NFSShareDefaultAccessReadOnly means: Allow read only access to the share for the hosts.
	NFSShareDefaultAccessReadOnly NFSShareDefaultAccessEnum = 1

	// NFSShareDefaultAccessReadWrite means: Allow read write access to the share for the hosts.
	NFSShareDefaultAccessReadWrite NFSShareDefaultAccessEnum = 2

	// NFSShareDefaultAccessRoot means: Allow read write root access to the share for the hosts.
	NFSShareDefaultAccessRoot NFSShareDefaultAccessEnum = 3

	// NFSShareDefaultAccessRoRoot means: Allow read only root access to the share for the hosts.
	NFSShareDefaultAccessRoRoot NFSShareDefaultAccessEnum = 4
)

type SeverityEnum int

const (

	// SeverityEMERGENCY means: Emergency
	SeverityEMERGENCY SeverityEnum = 0

	// SeverityALERT means: Alert
	SeverityALERT SeverityEnum = 1

	// SeverityCRITICAL means: Critical
	SeverityCRITICAL SeverityEnum = 2

	// SeverityERROR means: Error
	SeverityERROR SeverityEnum = 3

	// SeverityWARNING means: Warning
	SeverityWARNING SeverityEnum = 4

	// SeverityNOTICE means: Notice
	SeverityNOTICE SeverityEnum = 5

	// SeverityINFO means: information
	SeverityINFO SeverityEnum = 6

	// SeverityDEBUG means: debug
	SeverityDEBUG SeverityEnum = 7

	// SeverityOK means: OK
	SeverityOK SeverityEnum = 8
)

type SnapAccessLevelEnum int

const (

	// SnapAccessLevelReadOnly means: Allow read-only access to the snapshot for a host.
	SnapAccessLevelReadOnly SnapAccessLevelEnum = 0

	// SnapAccessLevelReadWrite means: Allow read/write access to the snapshot for a host.
	SnapAccessLevelReadWrite SnapAccessLevelEnum = 1

	// SnapAccessLevelReadOnlyPartial means: (Applies to consistency group snapshots only.) Indicates that host has read-only access to some individual snapshots in a consistency group snapshot. Do not use this value in Modify requests.
	SnapAccessLevelReadOnlyPartial SnapAccessLevelEnum = 2

	// SnapAccessLevelReadWritePartial means: (Applies to consistency group snapshots only.) Indicates that host has read/write access to some individual snapshots in a consistency group snapshot. Do not use this value in Modify requests.
	SnapAccessLevelReadWritePartial SnapAccessLevelEnum = 3

	// SnapAccessLevelMixed means: (Applies to consistency group snapshots only.) Indicates that host has read-only and read/write access to some individual snapshots in a consistency group snapshot. Do not use this value in Modify requests.
	SnapAccessLevelMixed SnapAccessLevelEnum = 4
)

type SnapCreatorTypeEnum int

const (

	// SnapCreatorTypeNone means: Not specified.
	SnapCreatorTypeNone SnapCreatorTypeEnum = 0

	// SnapCreatorTypeScheduled means: Created by a snapshot schedule.
	SnapCreatorTypeScheduled SnapCreatorTypeEnum = 1

	// SnapCreatorTypeUser_Custom means: Created by a user with a custom name.
	SnapCreatorTypeUser_Custom SnapCreatorTypeEnum = 2

	// SnapCreatorTypeUser_Default means: Created by a user with a default name.
	SnapCreatorTypeUser_Default SnapCreatorTypeEnum = 3

	// SnapCreatorTypeExternal_VSS means: Created by Windows Volume Shadow Copy Service (VSS) to obtain an application consistent snapshot.
	SnapCreatorTypeExternal_VSS SnapCreatorTypeEnum = 4

	// SnapCreatorTypeExternal_NDMP means: Created by an NDMP backup operation.
	SnapCreatorTypeExternal_NDMP SnapCreatorTypeEnum = 5

	// SnapCreatorTypeExternal_Restore means: Created as a backup snapshot before a snapshot restore.
	SnapCreatorTypeExternal_Restore SnapCreatorTypeEnum = 6

	// SnapCreatorTypeExternal_Replication_Manager means: Created by Replication Manager.
	SnapCreatorTypeExternal_Replication_Manager SnapCreatorTypeEnum = 8

	// SnapCreatorTypeReplication means: Created by a native replication operation.
	SnapCreatorTypeReplication SnapCreatorTypeEnum = 9

	// SnapCreatorTypeSnap_CLI means: Created inband by SnapCLI.
	SnapCreatorTypeSnap_CLI SnapCreatorTypeEnum = 11

	// SnapCreatorTypeAppSync means: Created by AppSync.
	SnapCreatorTypeAppSync SnapCreatorTypeEnum = 12
)

type SnapStateEnum int

const (

	// SnapStateReady means: The snaphot is operating normally.
	SnapStateReady SnapStateEnum = 2

	// SnapStateFaulted means: The storage pool that the snapshot belongs to is degraded.
	SnapStateFaulted SnapStateEnum = 3

	// SnapStateOffline means: The snapshot is not accessible possibly because the storage resource is not ready or the storage pool is full.
	SnapStateOffline SnapStateEnum = 6

	// SnapStateInvalid means: The snapshot has become invalid becauuse of a non recoverable error.
	SnapStateInvalid SnapStateEnum = 7

	// SnapStateInitializing means: The snapshot is being created.
	SnapStateInitializing SnapStateEnum = 8

	// SnapStateDestroying means: The snapshot is being deleted.
	SnapStateDestroying SnapStateEnum = 9
)
