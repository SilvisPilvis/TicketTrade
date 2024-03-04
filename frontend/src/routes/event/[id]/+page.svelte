<script>
// /** @type {import('./$types').PageData} */
export let data;
import axios from "axios";
import Cookies from 'js-cookie';
import { onMount } from "svelte";
import "iconify-icon"
// export let data;
let res, allReviews, failed, success, reviewRating, reviewComment;
onMount(async () => {
    try{
        const response = await axios.get(`http://localhost:8000/api/event/${data.data}`);
        res = response.data;
    }catch (e){
        console.error("Error:", e);
        failed = e.response.data.error;
    }
    try{
        const response = await axios.get(`http://localhost:8000/event/${data.data}/reviews`);
        allReviews = response.data;
    }catch (e){
        failed = e.response.data.error;
    }
})

function buyTicket(){
    window.location.replace(`/event/${data.data}/tickets`);
}

function postReview(){
    const config = {
            headers: { Authorization: `Bearer ${Cookies.get('token')}` }
    };
    if(reviewComment == "" || reviewComment == undefined ){
        failed = "Comment can't be empty.";
        return;
    }else{
        failed = "";
    }
    if(isNaN(reviewRating)){
        failed = "Rating must be a number in range from 1 to 5.";
    }
    if(reviewRating <= 0 || reviewRating > 5 || reviewRating == undefined){
        failed = "Rating must be in range from 1 to 5.";
        return;
    }else{
        failed = "";
    }
    console.log(Number(data.data));
    axios.post("http://127.0.0.1:8000/api/review", {
            eventId: Number(data.data),
            rating: reviewRating,
            comment: reviewComment
        }, config)
        .then(function (response) {
            // console.log(response);
            success = response.data;
            failed = "";
        })
        .catch(function (error) {
            console.log(error);
            failed = error.response.data.error;
            success = "";
        });
}

function deleteReview(id){
    const config = {
        headers: { Authorization: `Bearer ${Cookies.get('token')}` }
    };
    axios.delete(`http://127.0.0.1:8000/api/reviews/${id}`, config)
        .then(function (response) {
            console.log(response);
            success = response.data;
            failed = "";
        })
        .catch(function (error) {
            console.log(error);
            failed = error.response.data.error;
            success = "";
        });
}
</script>

<main class="flex col cen">
    <div class="margin-t">
        {#if res != "" && res != undefined}
            <!-- <img src={res.eventBanner} alt="" loading="lazy" class="banner"> -->
            <div class="flex row bg">
                <!-- <img class="cover" src={res.eventImage} alt="" loading="lazy"> -->
                <img src={res.eventBanner} alt="" loading="lazy" class="banner">
                <div class="description">
                    <h1>{res.eventName}</h1>
                    <div class="flex row gap1">
                        <div class="flex gap"><iconify-icon icon="bi:calendar-week" class="icon"></iconify-icon>{res.eventDate}</div>
                        <div class="flex gap"><iconify-icon icon="bi:geo-alt-fill" class="icon"></iconify-icon>{res.eventLocation}</div>
                    </div>
                    <div>{res.eventDescription}</div>
                    <div class="seats">Vietu skaits: {res.eventCapacity}</div>
                </div>
            </div>
        {/if}
        <div class="flex cen">
            <button on:click={buyTicket}>Buy Ticket</button>
        </div>
        <div class="flex cen row">
            <input type="text" bind:value={reviewComment} placeholder="Enter review here">
            <input type="number" name="" id="" min="1" max="5" step="1" bind:value={reviewRating}>
            <button on:click={postReview}>Post Review</button>
        </div>
        {#if failed != undefined && failed != ""}
            <p class="error">{failed}</p>
        {/if}
        {#if success != undefined && success != ""}
            <p class="success">{success}</p>
        {/if}
        {#if allReviews != "" && allReviews != undefined}
        <div class="flex col cen">
            {#each allReviews as review}
                <div class="review">
                    <div class="flex row between">
                        <p class="flex cen">{review.UserName}</p>
                        <button class="flex cen trash" on:click={() => deleteReview(review.Id)}><iconify-icon icon="mdi:trash-can"  style="color: var(--button-text)"></iconify-icon></button>
                    </div>
                    <p>{review.Rating}/5</p>
                    <p>{review.Comment}</p>
                </div>
            {/each}
        </div>
        {/if}
        
    </div>
</main>

<style>
.seats{
    font-weight: 700;
}
*{
    color: white;
}
img{
    border-radius: 0.4rem;
}
.gap{
    gap: 0.5rem;
    align-items: center;
}
.gap1{
    gap: 1rem;
    align-items: center;
}
.icon{
    /* margin: 0 1rem 0 1rem; */
}
.margin-t{
    margin-top: 5rem;
}
.between{
    justify-content: space-between;
}
.cover{
    margin-top: 1rem;
    margin-left: 1rem;
    border-radius: 0.4rem;
}
.banner{
    max-width: 100%;
    min-width: 80%;
/*    border-radius: 0.2rem;*/
}
.cover{
    max-width: min(20rem, 100%);
    border-radius: 0.3rem;
}
.description{
    width: 35rem;
}
.description>div{
    margin: 1rem;
}
.description>h1{
    margin: 1rem;
}
.bg{
    margin: 1rem 0 1rem 1rem;
    border-radius: 0.4rem;
    width: 57rem;
    padding-bottom: 1rem;
}
.review{
    margin: 1rem 0 0 1rem;
    background-color: var(--button-fill);
    width: 60rem;
    border-radius: 0.4rem;
    padding-bottom: 1rem;
}
.review>p{
    margin: 0.5rem 0 0 1rem;
}
.between>p{
    margin: 0.5rem 0 0 1rem;
}
.between>button{
    margin: 1rem 1rem;
}
.rev{
    margin: 1rem 1rem;
}
button{
    background-color: var(--fg);
    border: none;
    /* margin: 1rem 0 0 1rem; */
    padding: 1rem;
    border-radius: 0.4rem;
}
.trash{
    padding: 0.5rem;
}
input{
    margin: 1rem 1rem;
    padding: 1rem;
    border: none;
    border-radius: 0.4rem;
    background-color: var(--button-fill);
}
</style>