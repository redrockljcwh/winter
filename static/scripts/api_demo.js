import { Alexios } from './lib.js'; 

const result = document.querySelector('#result');

Alexios.get("/test").then(text => result.textContent = text)
