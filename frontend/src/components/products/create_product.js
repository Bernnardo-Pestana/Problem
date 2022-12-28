import axios from "axios";
import React, {useState } from "react";

function ProductCreate(){

    let [product, Setproduct] = useState({
        name:"",
        price:"",
    })


    async function create (){
       
         await axios.post(`http://127.0.0.1:3000/products`, product).then(res => console.log(res))
    }

    function handleCreate(evt){
        const value = evt.target.value;
        Setproduct({
            ...product,
            [evt.target.name]: value
          });
    }



    return(

        <div>
            <br/>
            <h2>Create Product</h2>
           <div>
                <label> Name</label>
                <input type="text" onChange={handleCreate} name="name" value={product.name}></input>
           </div>
           <div>
                <label> Price</label>
                <input type="text" onChange={handleCreate} name="price" value={product.price}></input>
            </div>
            <div>
                <button onClick={create}>Create</button>
            </div>
            <br/>
            <br/>
            <br/>
        </div>
    )
}


export default function ProductCreateView() {
    return (
      <section>
    
        <ProductCreate
        />

      </section>
    );
  }