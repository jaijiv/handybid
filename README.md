handybid
========

Sample social app written in go and mongo

- server design based on uncle bob clean architecture 
  http://blog.8thlight.com/uncle-bob/2012/08/13/the-clean-architecture.html

- package dependencies from inner circle going outwards :
  - domain
    data model definitions and interfaces
  - interfaces.repositories
    implementation of domain interfaces. DB is mongo
  - usecases
    application use case implementations. Also defines data types used across usecases and webcontrollers
  - infrastructure (same level as interfaces.webcontrollers)
  	DB connection specifics to mongo and other infra related code
  - interfaces.webcontrollers (same level as infrastructure)
  	REST interfaces. Defines interfaces implemented by usecases

  	main.go - usecases module is injected into webcontroller. user repository (interfaces.repositories.mongo implementation) is injected into usecases. This way each of the outer modules only provides implementations for interfaces provided by lower layer. This makes modules pluggable, easily mocked for testing and loosely coupled.