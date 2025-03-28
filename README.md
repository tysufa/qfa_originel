<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/tysufa/qfa">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">QFA</h3>

  <p align="center">
    The best programming language
    <br />
    <a href="https://github.com/tysufa/qfa"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/tysufa/qfa">View Demo</a>
    ·
    <a href="https://github.com/tysufa/qfa/issues">Report Bug</a>
    ·
    <a href="https://github.com/tysufa/qfa/issues">Request Feature</a>
  </p>
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
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

<!-- [![Product Name Screen Shot][product-screenshot]](https://example.com) -->

This is a programming language that I am creating for fun and to learn how a programming language really work.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With
* ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple example steps.

### Prerequisites

* Download go at [https://go.dev/doc/install](https://go.dev/doc/install)

### Installation
1. Clone the repo
   ```sh
   git clone https://github.com/tysufa/qfa.git
   ```
2. Start
   ```sh
   go run main.go
   ```

this will launch the REPL.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->

## Usage

### end of statement
every statement ends with a semicolon
```
statement;
```

### declarations
Use the keyword let
```
let variable = value;
```
### types
suported types are limited to integers and booleans at the moment. String and Arrays will be implemented later.
```
let variable1 = 10;
let variable2 = 123456789;
let variable3 = true;
let variable4 = false;
```
### functions
you can declare a function with the fn keyword and return a value with return
```
fn(x, y){
  return x + y;
};
```
and you can stock it like you would with a variable
```
let add = fn(x, y){
  return x + y;
};
```
### if statements
you can do an if else statement like you would with any language
```
if (10 == 9){
  expression;
}
else{
  expression;
}
```
<!-- _For more examples, please refer to the [Documentation](https://example.com)_ -->

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Lexer
	- Extand the lexer
		- [ ] Strings
		- [ ] floats
		- [ ] Arrays
- [ ] Parser
  - [x] Expressions
  - [x] If else
  - [ ] Functions
  - [ ] Extand with string arrays floats...

- [ ] Evaluator
  - [x] integers
  - [x] booleans
  - [x] If else
  - [ ] Functions

See the [open issues](https://github.com/tysufa/qfa/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Philemon PENOT - philemon.penot@gmail.com

Project Link: [https://github.com/tysufa/qfa](https://github.com/github_username/repo_name)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* []() Huge thanks to Thorsten Ball and his [book](https://interpreterbook.com/) wich has been the main source of inspiration for this project

* []() Go check as well the [teeny tiny compiler](https://austinhenley.com/blog/teenytinycompiler1.html) tutorial, it's a great way to begin the journey of compiling


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/tysufa/qfa.svg?style=for-the-badge
[contributors-url]: https://github.com/tysufa/qfa/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/tysufa/qfa.svg?style=for-the-badge
[forks-url]: https://github.com/tysufa/qfa/network/members
[stars-shield]: https://img.shields.io/github/stars/tysufa/qfa.svg?style=for-the-badge
[stars-url]: https://github.com/tysufa/qfa/stargazers
[issues-shield]: https://img.shields.io/github/issues/tysufa/qfa.svg?style=for-the-badge
[issues-url]: https://github.com/tysufa/qfa/issues
[license-shield]: https://img.shields.io/github/license/tysufa/qfa.svg?style=for-the-badge
[license-url]: https://github.com/tysufa/qfa/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: images/screenshot.png
[Next.js]: https://img.shields.io/badge/next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white
[Next-url]: https://nextjs.org/
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[Angular.io]: https://img.shields.io/badge/Angular-DD0031?style=for-the-badge&logo=angular&logoColor=white
[Angular-url]: https://angular.io/
[Svelte.dev]: https://img.shields.io/badge/Svelte-4A4A55?style=for-the-badge&logo=svelte&logoColor=FF3E00
[Svelte-url]: https://svelte.dev/
[Laravel.com]: https://img.shields.io/badge/Laravel-FF2D20?style=for-the-badge&logo=laravel&logoColor=white
[Laravel-url]: https://laravel.com
[Bootstrap.com]: https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white
[Bootstrap-url]: https://getbootstrap.com
[JQuery.com]: https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white
[JQuery-url]: https://jquery.com 
