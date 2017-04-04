# manifesto
Generate K8s manifest from templates
====================================
manifesto
==================
 A small script for generating Kubernetes manifests files using templates
Values can be defined on yaml file

Can be used for replacing, Tenant names, NameSpace, AppName etc
Can be customized to add any values you like to add

Benefits:

This makes easy to apply naming policies, best practices etc
Update policy/best practice changes quickly with less effort
Reduce time spent on creating or fixing K8 manifest during migration
Reduce the time for code review and approval for manifest files, as we can
follow a standard templates
Tenant/app specific values are store in values yaml

Install
==================

go build manifest.go

Usage
==================

CLI Options
==================
 values=values.yaml

 Accept the values for template. Like AppName, TenantName etc
 Default value is values.yaml

If you want to add new custom values add to struct Config in the code
Example: use values.yaml in this repo

tmplfile=svc.yaml

No default values
Example svc.yaml in this repo

ToDo
==================
Define and add more templates
Make this script work with Dockerfiles as well
CLI help
