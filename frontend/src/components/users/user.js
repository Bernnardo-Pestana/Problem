import axios from "axios";
import './user.css'
import UserCreateView from './create_user'
import React, {useState } from "react";

function User(){
    let [responseData, setResponseData] = React.useState([])

    function getData(e){
       axios.get('http://127.0.0.1:3000/users').then(res =>{
         setResponseData(res.data)
       })
     }
     const [permissions,setPermissions] = useState([]);


     function deleteSelected(){
        console.log("Irei deletar todos os selecionados....");
        console.log(permissions);
     }


     const handleCheck = (event) => {
        var permissions_array = [...permissions];
        if (event.target.checked) {
          permissions_array = [...permissions, event.target.value];
        } else {
          permissions_array.splice(permissions.indexOf(event.target.value), 1);
        }
        setPermissions(permissions_array);
      
      };


    const del = (id) =>{
      axios.delete(`http://127.0.0.1:3000/users/${id}`)
      .then(response => console.log(response));
    }

     /* 
      function edit(id){
        axios.put(`http://127.0.0.1:3000/users/${id}`, )
        .then(response => console.log(response));
      } 

     
     */
      

    


    return(
        <div className="main">
          <div>
          <h2> Index of User</h2>
            <button  onClick={(e) => getData(e)}> Click to Load</button>
            <table >
                <thead >
                <tr>
                    <th></th>
                    <th className="table-header">Email</th>
                    <th className="table-header">First Name</th>
                    <th className="table-header">Last Name</th>
                    <th className="table-header">User</th>
                    <th className="table-header">Options</th>
                </tr>
                </thead>
                <tbody>
                    {
                       responseData.map((user,index) =>{
                        return (
                            <tr key={index}>
                                <th> <input type="checkbox" value={user.id} onChange={handleCheck}/></th>
                                <th className="table-content">{user.email }</th>
                                <th className="table-content"> {user.first_name }</th>
                                <th className="table-content">{user.last_name } </th>
                                <th className="table-content">{user.type }</th>
                                <th > <button >Edit</button> <button onClick={()=> del(user.id)}>Delete</button></th>
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

          </div>
        </div>
    )
 
}



export default function Users() {
    return (
      <section>
        <UserCreateView />
        <User />
        

      </section>
    );
  }