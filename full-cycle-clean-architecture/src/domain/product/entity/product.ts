import ProductAbstract from "./product-abstract";
import ProductInterface from "./product.interface";

export default class Product extends ProductAbstract {

  constructor(id: string, name: string, price: number) {
    super(id, name, price)
  }


}
