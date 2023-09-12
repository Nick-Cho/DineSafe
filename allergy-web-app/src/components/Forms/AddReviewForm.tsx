import React, {useState} from 'react'
import axios from "axios";

function AddReviewForm() {
    const [allergy, setAllergy] = useState<string>(""); // Tracks the allergy selected by the user in the review
    const [item, setItem] = useState<string>(""); // Tracks the item that the user has had an allergic reaction with 
    return (
        <div className="mt-5 w-full bg-white p-5 rounded-lg">
            <form onSubmit={()=>{}}>
                <select 
                    name="allergies" 
                    className="mb-5 px-3 py-2 font-uber bg-gray rounded-lg"
                    onChange={(e)=>{setAllergy(e.target.value)}}
                >
                    <option value="nuts">Nuts</option>
                    <option value="peanuts">Peanuts</option>
                    <option value="sesame">Sesame</option>
                    <option value="egg">Eggs</option>
                    <option value="milk">Milk</option>
                    <option value="shellfish">Shellfish</option>
                    <option value="fish">Fish</option>
                    <option value="soy">Soy</option>
                    <option value="gluten">Gluten</option>
                </select>
                <input
                    type="text"
                    className="font-uber rounded-lg py-3 px-4 w-full bg-gray"
                    placeholder="Which item should people look out for?"
                    onChange={(e)=>{setItem(e.target.value)}}
                />
                <button className="bg-black text-white font-uber rounded-lg font-medium px-2 py-1 mt-3 cursor-pointer">
                    Add Review
                </button>
            </form>                   
        </div>
        )
    }

export default AddReviewForm