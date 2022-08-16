# api-crud-mysql

## Features

- [x] Consume API
- [x] Store into a Mysql database
- [x] CRUD 
- [ ] Use an interface for CRUD operations

## What to expect from this project: 

- this project consume a API (http://pokeapi.co/api/v2/pokedex/kanto/);
- put all information from the API into the database (Mysql);
- and allow the user to access the information, create new pokemons, update the existing ones and delete then;

## Packages:

To create this CRUD (Create, Read, Update and Delete) project we use the following packages

- ``gorm.io/gorm``
- ``InteliJ IDEA``
- ``Paradigma de orientação a objetos``

1) gorm.io/gorm
>used to conect the Mysql database
2) github.com/spf13/viper
>used to easy config everything in the project
3) github.com/gorilla/mux
> used to defines and use routes
4) net/http
> used to consume the Pokemon API

## Project Structure:

