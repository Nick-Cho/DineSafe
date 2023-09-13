import React, {useState} from 'react'
import env from "react-dotenv"
import axios from "axios";

type Props = {
    address: string
    setAddReview : React.Dispatch<React.SetStateAction<boolean>>
}

function AddReviewForm({address, setAddReview}: Props) {
    const [allergy, setAllergy] = useState<string>(""); // Tracks the allergy selected by the user in the review
    const [item, setItem] = useState<string>(""); // Tracks the item that the user has had an allergic reaction with 
    const handleSubmit = async (e: React.SyntheticEvent<EventTarget>) => {
        e.preventDefault();
        try {
            const requestBody = JSON.stringify({
                review: item,
                allergy: allergy,
                street_address: address
            })
            // console.log("Allergy selection: ", allergy);
            const response = await axios.post(`${env.API_URL}/addReview`,requestBody);
            if (response.status === 202) {
                setAddReview(false);
            }
            console.log("Response from addReview endpoint: ", response);
        } catch (err) {
            console.log("Error from addReview call: ", err);
        }
    }

    return (
        <div className="mt-5 w-full bg-white p-5 rounded-lg">
            <form onSubmit={(e)=>{handleSubmit(e)}}>
                <select 
                    name="allergies" 
                    value={allergy}
                    className="mb-5 px-3 py-2 font-uber bg-gray rounded-lg"
                    onChange={(e)=>{setAllergy(e.target.value)}}
                >
                    <option selected value="nuts">Nuts</option>
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
                <button 
                    className="bg-black text-white font-uber rounded-lg font-medium px-2 py-1 mt-3 cursor-pointer"    
                >
                    Add Review
                </button>
            </form>                   
        </div>
        )
    }

export default AddReviewForm