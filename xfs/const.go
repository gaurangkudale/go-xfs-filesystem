package xfs

const (
	BMBT_EXNTFLAG_BITLEN = 1
	INODEV3_SIZE         = 176
	INODE_SIZE           = 96
	LEAF_ENTRY_SIZE      = 8

	XFS_DIR2_DATA_FD_COUNT  = 3
	XFS_DIR2_DATA_FREE_TAG  = 0xffff
	XFS_DIR2_DATA_ALIGN_LOG = 3

	XFS_SB_MAGIC         = 0x58465342
	XFS_AGF_MAGIC        = 0x58414746
	XFS_AGI_MAGIC        = 0x58414749
	XFS_AGFL_MAGIC       = 0x5841464c
	XFS_DINODE_MAGIC     = 0x494e
	XFS_DQUOT_MAGIC      = 0x4451
	XFS_SYMLINK_MAGIC    = 0x58534c4d
	XFS_ABTB_MAGIC       = 0x41425442
	XFS_ABTB_CRC_MAGIC   = 0x41423342
	XFS_ABTC_MAGIC       = 0x41425443
	XFS_ABTC_CRC_MAGIC   = 0x41423343
	XFS_IBT_MAGIC        = 0x49414254
	XFS_IBT_CRC_MAGIC    = 0x49414233
	XFS_FIBT_MAGIC       = 0x46494254
	XFS_FIBT_CRC_MAGIC   = 0x46494233
	XFS_BMAP_MAGICa      = 0x424d4150
	XFS_BMAP_CRC_MAGIC   = 0x424d4133
	XFS_DA_NODE_MAGIC    = 0xfebe
	XFS_DA3_NODE_MAGIC   = 0x3ebe
	XFS_DIR2_BLOCK_MAGIC = 0x58443242
	XFS_DIR3_BLOCK_MAGIC = 0x58444233 // XDB3 Block Directory Magic number
	XFS_DIR2_DATA_MAGIC  = 0x58443244
	XFS_DIR3_DATA_MAGIC  = 0x58444433 // XDD3 Leaf Directory Magic number
	XFS_DIR2_LEAF1_MAGIC = 0xd2f1
	XFS_DIR3_LEAF1_MAGIC = 0x3df1
	XFS_DIR2_LEAFN_MAGIC = 0xd2ff
	XFS_DIR3_LEAFN_MAGIC = 0x3dff
	XFS_DIR2_FREE_MAGIC  = 0x58443246
	XFS_DIR3_FREE_MAGIC  = 0x58444633
	XFS_ATTR_LEAF_MAGIC  = 0xfbee
	XFS_ATTR3_LEAF_MAGIC = 0x3bee
	XFS_ATTR3_RMT_MAGIC  = 0x5841524d
	XFS_RMAP_CRC_MAGIC   = 0x524d4233
	XFS_RTRMAP_CRC_MAGIC = 0x4d415052
	XFS_REFC_CRC_MAGIC   = 0x52334643
	XFS_MD_MAGIC         = 0x5846534d
)

const (
	XFS_DIR2_DATA_SPACE = iota
	XFS_DIR2_LEAF_SPACE
	XFS_DIR2_FREE_SPACE
)

const (
	// typedef enum xfs_dinode_fmt
	XFS_DINODE_FMT_DEV = iota
	XFS_DINODE_FMT_LOCAL
	XFS_DINODE_FMT_EXTENTS
	XFS_DINODE_FMT_BTREE
	XFS_DINODE_FMT_UUID
	XFS_DINODE_FMT_RMAP
)
