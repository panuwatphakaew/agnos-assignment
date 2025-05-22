package hospital

type Hospital struct {
	ID   int    `json:"id" bun:"id,pk autoincrement"`
	Name string `json:"name" bun:"name" bun:",unique,notnull"`
}
