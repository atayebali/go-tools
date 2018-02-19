Gogo Github
===========
Create Branches and Generate PRs for mutliple repos. 

You will need a Github Token as an ENV var to generate PRs via this app.  See link below to obtain one.

[Personal Dev Token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/)

Token should be saved in GITHUB_TOKEN for it to work.

Syntax: 
=======

### SETUP:
This will clone all the FE apps into ~/FRONT_END_APPS:
```
pr-spin -setup y
```

### BUILD A FEATURE:
 This takes a branch name and keys to the repos you need initialized
```    
pr-spin -make-feature -branch feature/HTW-1111-fix-table -repos TW,FB,CT,SH
```

### PUSH A FEATURE
```
pr-spin -push-feature -branch feature/HTW-1111-fix-table -repos TW,FB,CT,SH	
```


Commands:
=========
	-setup:  
		Clones all the Front End Repos into ~/FRONT_END_APPS

	-make-feature:
		Creates a branch off master and pushes to remote. Updates package.json for Shell App
		makes a new yarn.lock and creates a PR.
		
		-branches:
			The branch that will be cut from master
			
		-repos:
			List of CSV keys corresponding the apps you need updating. 

	-push-feature:
		Commits changes on the branches you are working on and pushes to GH 
		creates	PR with master as base for all apps. 

		-branches:
			The branch that will be cut from master
			
		-repos:
			List of CSV keys corresponding the apps you need updating. 