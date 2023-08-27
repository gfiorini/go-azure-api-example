Pre-requisites:

you must have installed the following tools in your local machine 

- .net 6.0 SDK (https://dotnet.microsoft.com/en-us/download/dotnet/6.0)
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
http://localhost:7071/api/HttpExample