# Usage

Step on BuggyBuddy directory:

`export ROOT=$(PWD)`

Create and apply the manifests:

`kustomization build $ROOT | kubectl apply -f -n foo`

## TODO:

* Create a patch to select image repo