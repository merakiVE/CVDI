# Reference from API CVDI

## Neuron
> **Model:** NeuronModel<br/>
> **Source model:** src/models/neuron.go<br/>
> **Controller:** src/controllers/neuron.go<br/>
> **Endpoint:** /neurons

Small Description

## Methods

|Enpoint                |Method         |Description                         |
|----------------|-------------|----------------|
|/|GET          |Get all neuros|
|/{neuronKey:string}|GET|Get specific neuron with key|
|/{neuronKey:string}/actions|GET| Get all action from specific neuron key|
|/|POST|Create new neuron|
|/{neuronKey:string}/actions|POST|Create or add new action to neuron|

<br/>

## Auth
> **Model:** UserModel<br/>
> **Source model:** src/models/user.go<br/>
> **Controller:** src/controllers/auth.go<br/>
> **Endpoint:** /auth

Small Description

## Methods

|Enpoint                |Method         |Description                         |
|----------------|-------------|----------------|
|/login|POST          |Login user in system|

<br/>

## Users
> **Model:** UserModel<br/>
> **Source model:** src/models/user.go<br/>
> **Controller:** src/controllers/user.go<br/>
> **Endpoint:** /users

Small Description

## Methods

|Enpoint                |Method         |Description                         |
|----------------|-------------|----------------|
|/{userKey:string}|GET|Get specific user with key|
|/|GET|Get all user|
|/|POST|Create new user|
