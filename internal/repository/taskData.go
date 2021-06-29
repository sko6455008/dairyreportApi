package repository

type TaskProject string

const (
	TaskProject0 TaskProject = "その他"

	TaskProject1 TaskProject = "おしるこ"

	TaskProject2 TaskProject = "東京メトロ"

	TaskProject3 TaskProject = "ふれあいリズムダンス"

	TaskProject4 TaskProject = "相模屋美術店"

	TaskProjectGC TaskProject = "GC"

	TaskProjectItemstore TaskProject = "itemstore"

	TaskProjectJRA TaskProject = "JRA"

	TaskProjectMtg TaskProject = "社内(mtg/調査など)"
)

type TaskData struct {
	Hour    float32
	Project TaskProject
}
