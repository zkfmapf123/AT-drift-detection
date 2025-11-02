package usecase

type AtlantisRequest struct {
	GithubToken   string
	GithubRepoRef string

	AtlantisURL        string
	AtlantisToken      string
	AtlantisRepository string
	AtlantisConfigFile string
}

type APIHealthResponse struct {
	Status string `json:"status"`
}

type APIPlanBodyParams struct {
	Repository string             `json:"repository"`
	Ref        string             `json:"ref"`
	Type       string             `json:"type"` // Gtihub
	Paths      []APIPlanBodyPaths `json:"paths"`
}

type APIPlanBodyPaths struct {
	Directory   string `json:"Directory"`
	Workspace   string `json:"Workspace,omitempty"`   // default
	ProjectName string `json:"ProjectName,omitempty"` // 중복 프로젝트 구분용
}

type APIPlanResponse struct {
	// Error          interface{}     `json:"Error"`
	// Failure        string          `json:"Failure"`
	ProjectResults []struct {
		PlanSuccess struct {
			TerraformOutput string `json:"TerraformOutput"`
		} `json:"PlanSuccess"`
	} `json:"ProjectResults"`
	// PlansDeleted   bool            `json:"PlansDeleted"`
}

type ProjectResult struct {
	// Command            int          `json:"Command"`
	// SubCommand         string       `json:"SubCommand"`
	// RepoRelDir         string       `json:"RepoRelDir"`
	// Workspace          string       `json:"Workspace"`
	// Error              interface{}  `json:"Error"`
	// Failure            string       `json:"Failure"`
	PlanSuccess *PlanSuccess `json:"PlanSuccess"`
	// PolicyCheckSuccess interface{}  `json:"PolicyCheckSuccess"`
	// ApplySuccess       string       `json:"ApplySuccess"`
	// VersionSuccess     string       `json:"VersionSuccess"`
	// ProjectName        string       `json:"ProjectName"`
}

type PlanSuccess struct {
	TerraformOutput string `json:"TerraformOutput"`
	// LockURL         string `json:"LockURL"`
	// RePlanCmd       string `json:"RePlanCmd"`
	// ApplyCmd        string `json:"ApplyCmd"`
	// HasDiverged     bool   `json:"HasDiverged"`
	// MergedAgain     bool   `json:"MergedAgain"`
}

/*
@examples
## example
version: 3
projects:
  - name: ec2
    dir: examples/ec2
    workflow: terraform
  - name: sg
    dir: examples/sg
    workflow: terraform

workflows:

	terraform:
	  plan:
	    steps:
	      - init:
	          extra_args: ["-upgrade", "-reconfigure"]
	      - env:
	          name: TERRAGRUNT_TFPATH
	          command: 'echo "terraform${ATLANTIS_TERRAFORM_VERSION}"'
	      - env:
	          name: TF_IN_AUTOMATION
	          value: "true"
	      - run:
	          command: terraform plan -input=false -out=$PLANFILE
	          output: strip_refreshing
	  apply:
	    steps:
	      - env:
	          name: TERRAGRUNT_TFPATH
	          command: 'echo "terraform${ATLANTIS_TERRAFORM_VERSION}"'
	      - env:
	          name: TF_IN_AUTOMATION
	          value: "true"
	      - run: terraform apply $PLANFILE
*/
type AtlantisConfigParams struct {
	Version  string `yaml:"version"`
	Projects []struct {
		Name     string `yaml:"name"`
		Dir      string `yaml:"dir"`
		Workflow string `yaml:"workflow"`
	} `yaml:"projects"`
}
