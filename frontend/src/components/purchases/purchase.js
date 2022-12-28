import axios from "axios";
import React, {useState } from "react";

function Purchase(){

    let [responseData, setResponseData] = React.useState([]);

    

    function getData(e){
       axios.get('http://127.0.0.1:3000/purchases').then(res =>{
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
        axios.get(`http://127.0.0.1:3000/purchases/${id}`).then(res =>{
           
    
            SetEditProduct(res.data)
           })
    }

    const del = (id) =>{
        console.log("Deletar ",id)
    }

    function handleChangeEdit(evt) {
        const value = evt.target.value;
        SetEditProduct({
          ...editProduct,
          [evt.target.name]: value
        });
      }

    function edite(id){
        axios.put(`http://127.0.0.1:3000/purchases/${id}`, editProduct)
        .then(response => console.log(response));


    }

    async function create (){
        await axios.post(`http://127.0.0.1:3000/purchases`, {name:"Pencil2",price:1.55}).then(res => console.log(res))
    }

    return(
        <div className="main">
            <div>
            <button  onClick={(e) => getData(e)}> Click to Load data...</button>
            <table>
                <thead>
                <tr>
                    <th className="table-header"></th>
                    <th className="table-header">Id_User</th>
                    <th className="table-header">Id_Product</th>
                    <th className="table-header">Options</th>
                </tr>
                </thead>
                <tbody>
                    {
                       responseData.map((product,index) =>{
                        return (
                            <tr key={index}>
                                <th > <input type="checkbox" value={product.id} onChange={handleCheck}/></th>
                                <th className="table-content">{product.id_user }</th>
                                <th className="table-content"> {product.id_product}</th>
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



export default function Purchases() {
    return (
      <section>
    
        <Purchase />

      </section>
    );
  }