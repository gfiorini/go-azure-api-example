Pre-requisites:

you must have installed the following tools in your local machine 

- at least .net 6.0 SDK (https://dotnet.microsoft.com/en-us/download/dotnet/6.0)
- azure functions tool (https://github.com/Azure/azure-functions-core-tools/blob/v4.x/README.md#windows)
- copy local.settings.json.template to local.settings.json (local.settings.json is in .gitignore because it can contain azure secrets and must not be versioned)

---------------------------

##### windows #####

# local build
winbuild.cmd (create go executable, handler.exe)

# local run function app (start local function app environment)
run.cmd

-------------------------------------------------
##### linux #####

# local build (create go executable, handler)
./build.sh

# local run function app (start local function app environment)
./run.sh

-------------------------------------------------

if everything is fine you should be able to call the function:
http://localhost:7071/api/albums


#### CI/CD ######

# Create Azure resource group (win command prompt)

set rg_name=rg-board-prod-001
set subscription_id=b844422f-15e5-4938-bd3c-6d24ea44c7d6
set location=germanywestcentral
set function_name=func-board-001

az login 
az group create -n %rg_name% -l %location%
az ad sp create-for-rbac --name "%rg_name%" --role Contributor --scopes /subscriptions/%subscription_id%/resourceGroups/%rg_name% --sdk-auth

# add in github the SECRETS: (from github.com/<user_name>/<project_name> go to settings -> secrets and variables -> actions -> new repository secret)

AZURE_RBAC_CREDENTIALS with the values retrieved from previous command
AZURE_RG=%rg_name%
AZURE_SUBSCRIPTION=%subscription_id%

# add in gitub the VARIABLE (tab variable -> new repository variable)

AZURE_FUNCTIONAPP_NAME=%function_name%
