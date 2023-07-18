import ProductAbstract from "./product-abstract";

export default class ProductB extends ProductAbstract {

  constructor(id: string, name: string, price: number) {
    super(id, name, price)
  }
}