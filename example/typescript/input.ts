enum Fruits {
  Apple = 1,
  Banana,
  Cherry,
  Pear,
}

class Person {
  #name: string;

  constructor(name: string) {
    this.#name = name;
  }

  get name(): string {
    return this.#name;
  }

  set name(value: string) {
    this.#name = value;
  }

  sayName() {
    console.log(`My name is ${this.#name}`);
  }
}

export { Fruits, Person };

// side effects
function init() {
  console.log("Hello World!");
}
window.onload = () => {
  init();
};
