# Easy Clone
Clone Organization Repositories easily with one go.

## Getting Started
Make sure that you have created Personal Access token which has the required permissions to clone.

### Prerequisites

* [Create Access Token from here](https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token)
* (Skip if not required)Enable SSO integration with your Org SSO provider
* Careful: Copy/Store Token at your local

### Installing
* [Golang is installed](https://golang.org/dl/)
* Clone this Repo
* Open .env in root folder

*   (Important) About Environment Variables

    Name | #Usage |
    --- | --- |
    GITHUB_ORG_API | Github API link (Do not change, unless required) |
    ACCESS TOKEN  | (Mandatory)Replace with your Access Token here, obtained from initial steps |
    GITHUB_USER_NAME | Provide your Github Username |
    DESTINATION_DIRECTORY | (Mandatory)Give path of the Directory for cloning repos |
    ORGANIZATION | Provide Github Organization Name |
    IGNORE_REPOS | (Optional)Skip Repos which are not required |
    IGNORE_FORK_REPOS | Skip Cloning Forked Repos |
    | |
* Update .env file with your values
* Open terminal from this repo directory
    ````
    Run >> go run main.go
    ````

* Watch Output for Status
* <b>Process may take time, (Depends on your network speed and total repos in your org)</b>

### Sample Output
<img alt="SampleOutout" src="https://user-images.githubusercontent.com/32619461/105634730-96499c80-5e85-11eb-911e-bfd0fbe57e30.png" width="450">

### Coding Format
This project follows a DI approach,
wherein implementation will be abstracted.

## Built With
* [Golang](https://golang.org/dl/) - Tool/Language used

## Contributing
Open to fork,clone,issues

## Authors
* **Chaitanya Kumar** 

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

