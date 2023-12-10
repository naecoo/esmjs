export class Person {
  #name;

  constructor(name) {
    this.#name = name;
  }

  get name() {
    return this.#name;
  }

  set name(value) {
    this.#name = value;
  }

  sayName() {
    console.log(`My name is ${this.#name}`);
  }
}

console.log("Hello World!");
