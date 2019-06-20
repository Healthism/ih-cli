package cmd

/**/
var RELEASE string
var CLUSTER string
/**/

const VERSION = "1.0.3"
const EDITOR = "vim"

const CLUSTER_TEMPLATE = "gke_%s_northamerica-northeast1-a_%s"
const COFING_PATH = "/usr/local/lib/ih"
const JOB_PATH = "/usr/local/lib/ih/values.yaml"
const IMAGE_PATH = "/usr/local/lib/ih/substitutions/_APP_IMAGE_URL"

/** ROOT **/
const IH = "ih"
const ROOT_DESCRIPTION_SHORT = "InputHealth Command Line Interface."
const ROOT_DESCRIPTION_LONG = "InputHealth Command Line Interface."

/** RUN **/
const RUN = "run"
const RUN_DESCRIPTION_SHORT = "InputHealth Command Line Interface."
const RUN_DESCRIPTION_LONG = "InputHealth Command Line Interface."

/** UPDATE **/
const UPDATE = "update"
const UPDATE_DESCRIPTION_SHORT = "InputHealth Command Line Interface."
const UPDATE_DESCRIPTION_LONG = "InputHealth Command Line Interface."
