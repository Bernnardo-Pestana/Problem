import axios from "axios";
import React, {useState } from "react";
import ProductCreateView from './create_product'
import './product.css'

function Product(){

    let [responseData, setResponseData] = React.useState([]);

    

    function getData(e){
       axios.get('http://127.0.0.1:3000/products').then(res =>{
        console.log(res.data)
         setResponseData(res.data)
       })
     }
     const [permissions,setPermissions] = useState([]);


     function deleteSelected(){
        console.log("Irei deletar todos os selecionados....");
        console.log(permissions);
     }

     let [editProduct,SetEditProduct] = useState();


     const handleCheck = (event) => {
        var permissions_array = [...permissions];
        if (event.target.checked) {
          permissions_array = [...permissions, event.target.value];
        } else {
          permissions_array.splice(permissions.indexOf(event.target.value), 1);
        }
        setPermissions(permissions_array);
      
      };

    const edit = (id) =>{
        axios.get(`http://127.0.0.1:3000/product/${id}`).then(res =>{
           
    
            SetEditProduct(res.data)
           })
    }

   

    function handleChangeEdit(evt) {
        const value = evt.target.value;
        SetEditProduct({
          ...editProduct,
          [evt.target.name]: value
        });
      }
      const del = (id) =>{
        console.log("Deletar ",id)
    }
    function edite(id){
        axios.put(`http://127.0.0.1:3000/product/${id}`, editProduct)
        .then(response => console.log(response));


    }

    async function create (){
        await axios.post(`http://127.0.0.1:3000/products`, {name:"Pencil2",price:1.55}).then(res => console.log(res))
    }

    return(
       <div className="main">
         <div>
            <button  onClick={(e) => getData(e)} className="button"> Click to Load data...</button>
            <table>
                <thead>
                <tr>
                    <th></th>
                    <th className="table-header">Name</th>
                    <th className="table-header">Price</th>
                    <th className="table-header">Options</th>
                </tr>
                </thead>
                <tbody>
                    {
                       responseData.map((product,index) =>{
                        return (
                            <tr key={index}>
                                <th> <input type="checkbox" value={product.id} onChange={handleCheck}/></th>
                                <th className="table-content">{product.name }</th>
                                <th className="table-content"> US$ {product.price }</th>
                                <th> <button onClick={()=> edit(product.id)}>Edit</button> <button onClick={()=> del(product.id)}>Delete</button></th>
                            </tr>
                        )
                       }) 
                    }
                </tbody>


            </table>
                    {responseData.length > 0 &&
                        <div>
                            <button onClick={(e) => deleteSelected()}>
                                Click to Delete All selected
                            </button>
                        </div>
                    }


                    { 
                        editProduct != null && 
                        <div>
                            <br/>
                            <br/>
                            <input  type="text" onChange={handleChangeEdit} value={editProduct.name}   name="name"/>
                            <input  type="text"  onChange={handleChangeEdit} value={editProduct.price} name="price"  />

                            <button onClick={()=> edite(editProduct.id)}>Edit</button>
                        </div>
                    
                    
                    }

        </div>
       </div>
    )
 
}



export default function Products() {
    return (
      <section>
    
        <ProductCreateView />
        <Product />

      </section>
    );
  }