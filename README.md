# GoCountry-App
## About
Console Application made with Go and Cobra Framework which returns countries information and more!

## Commands
### Country
Writes personalized countries information in console or on an csv file.
#### Flags
- skip: skip values from response. Default: 0. Short: s
- limit: limit number of values to response. Default: 5. Short: l
- name: get countries by name. Intersection with region if both flags are defined. Short: n
- region: get countries by region. Intersection with name if both flags are defined. Short: r
- to-csv: boolean to define if you want countries' information on a csv file. Default: true. Short: t
- file-name: personalize output csvfile name. Default: countries. Short: f
- console: boolean to define if you want countries' information just to be written on console. Default: false.Short: c
### Write
Writes absolutely all countries on a csv file.
#### Flags
- file-name: personalize output csvfile name. Default: countries. Short: f

## Future of the project
- Improve file-handling repository
- Avoid overwriting info and eliminate duplicated information
- Add Weather functionality
- Dockerize project

<img align="right" width=250 height=250 src="https://github.com/ashleymcnamara/gophers/blob/master/DockerGopher.png"/>
