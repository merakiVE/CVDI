# Reference from API CVDI

## Neuron
> **Model:** Neuron<space><space>
> **Source model:** src/models/neuron.go<space><space>
> **Endpoint:** /neurons<space><space>

Small Description

## Methods

|Enpoint                |Method         |Description                         |
|----------------|-------------|----------------|
|/|GET          |Get all neuros|
|/{neuronKey:string}|GET|Get specific neuron with key|
|/{neuronKey:string}/actions|GET| Get all action from specific neuron key|
|/|POST|Create new neuron|
|/{neuronKey:string}/actions|POST|Create or add new action to neuron|

