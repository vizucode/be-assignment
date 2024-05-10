<br />
<div align="center">

<h3 align="center">ConcentrateAI Test</h3>

<p align="center">
    Simple banking application that have feature transfer and withdraw, creating a account and secure authorization
</div>

<!-- TABLE OF CONTENTS -->

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#sequencediagram">Sequence Diagram</a></li>
    <li><a href="#erd">ERD</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

Simple banking application that have feature transfer and withdraw, creating a account and secure authorization

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

* Golang
* Nginx
* Gin Framework
* Docker
* Supabase
* PostgreSQL

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

To get started installing and using this app, you must first follow these steps. It's important as the app might not run otherwise!

### Prerequisites

- Register and log in to Supabase to obtain the API key and URI. This is necessary to connect Supabase to our app.
- Make sure you have Docker installed on your machine.
- Copy the `.env-example` file and create a `.env` file. Fill in all the required variables.
- We use port 5432 for PostgreSQL and port 80 for Nginx, so make sure there are no conflicting ports.

### Installation

#### Starting the docker compose

```bash
make dev up
```

##### Stop the docker compose

```bash
make dev down
```

<!-- USAGE EXAMPLES -->

## Usage

For swagger UI documentation

_For more documentation, please refer to the [APIDOG](https://tpdlhcvjhe.apidog.io)_


<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- SEQUENCE DIAGRAM -->

## Sequence Diagram

### Auth

![Auth Sequence Diagram](https://firebasestorage.googleapis.com/v0/b/personal-website-1d263.appspot.com/o/concentrateAI%2Fsequence-auth.png?alt=media&token=ada33fd5-0f65-4aa8-bf5b-f643193c689b)

### Create Account

![Create Account](https://firebasestorage.googleapis.com/v0/b/personal-website-1d263.appspot.com/o/concentrateAI%2Fcreate-account.png?alt=media&token=b561a7e4-f5f2-4524-9498-fba3282716b8)

### Send Transfer

![Send Transfer](https://firebasestorage.googleapis.com/v0/b/personal-website-1d263.appspot.com/o/concentrateAI%2Fsend-transfer.png?alt=media&token=07521e5d-9832-461d-91d2-e9090f8fceff)

### Withdraw Transfer

![Withdraw Transfer](https://firebasestorage.googleapis.com/v0/b/personal-website-1d263.appspot.com/o/concentrateAI%2Fwithdraw.png?alt=media&token=534b3019-5af0-45cc-817d-daecc2c2479c)

### Get Transaction

![Get Transaction](https://firebasestorage.googleapis.com/v0/b/personal-website-1d263.appspot.com/o/concentrateAI%2Fget-all-transaction.png?alt=media&token=9710ca99-f72f-42a0-931d-ca3cab71ea2a)

### Get Accounts

![Get Account](https://firebasestorage.googleapis.com/v0/b/personal-website-1d263.appspot.com/o/concentrateAI%2Fget-accounts.png?alt=media&token=886f76e9-8d54-4b47-8548-57b814fe6db7)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ERD -->

## ERD

![Get Account](https://firebasestorage.googleapis.com/v0/b/personal-website-1d263.appspot.com/o/concentrateAI%2Fconcreate-ai-test.png?alt=media&token=f94998d9-23c4-434c-9308-af49a0336c07)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

Hafiz Iqbal Sahrunizar - havisiqbalsyah@gmail.com

Github Profile: [https://github.com/vizucode](https://github.com/vizucode)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
