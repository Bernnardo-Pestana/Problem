function PurchaseCreate(){

    let [user, Setuser] = useState({
        email: "",
        first_name:"",
        last_name:"",
        password:"",
        type: "",
    })


    async function create (){
       
         await axios.post(`http://127.0.0.1:3000/users`, user).then(res => console.log(res))
    }

    function handleCreate(evt){
        const value = evt.target.value;
        Setuser({
            ...user,
            [evt.target.name]: value
          });
    }



    return(

        <div>
            <br/>
            <h2>Purchase User</h2>
           <div>
                <label> Email</label>
                <input type="text" onChange={handleCreate} name="email" value={user.email}></input>
           </div>
           <div>
                <label> First Name</label>
                <input type="text" onChange={handleCreate} name="first_name" value={user.first_name}></input>
            </div>
            <div>
                <label> Last Name</label>
                <input type="text"  onChange={handleCreate} name="last_name" value={user.last_name}></input>
            </div>
            <div>
                <label> Type : Customer - Seller</label>
                <input type="text" onChange={handleCreate} name="type" value={user.type}></input>
            </div>
            <div>
                <label> Password</label>
                <input type="password" onChange={handleCreate} name="password" value={user.password} ></input>
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


export default function PurchaseCreateView() {
    return (
      <section>
    
        <PurchaseCreate
        />

      </section>
    );
  }