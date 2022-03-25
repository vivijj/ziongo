package record

type Frstr string

type BlockStatus string

var (
	// Pending when the full block was generated.
	Pending = BlockStatus("pending")
	// Proved when the proof for the block is ready.
	Proved = BlockStatus("proved")
	// Committed when the block is submitted to L1 chain but not mined yet.
	Committed = BlockStatus("committed")
	// Verified when the block has been finalized(mined on L1 for required confirmation)
	Verified = BlockStatus("verified")
)

type StorageBlock struct {
	BlockNumber   int         `json:"block_number"`
	BlockSize     int         `json:"block_size"`
	NewStateRoot  Frstr       `json:"new_state_root"`
	Status        BlockStatus `json:"status,omitempty"`
	ProvedAt      int         `json:"proved_at,omitempty"`
	CommittedAt   int         `json:"committed_at,omitempty"`
	CommittedHash string      `json:"committed_hash,omitempty"`
	VerifiedAt    int         `json:"verified_at,omitempty"`
}
