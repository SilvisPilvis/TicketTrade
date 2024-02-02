<script>
// /** @type {import('./$types').PageData} */
export let data;
import axios from "axios";
import Cookies from 'js-cookie';
import { onMount } from "svelte";
import "iconify-icon"
// export let data;
let res, allReviews, failed, success, reviewRating, reviewComment, userId = 1;
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

function postReview(){
    const config = {
            headers: { Authorization: `Bearer ${Cookies.get('token')}` }
        };
    if(reviewRating <= 0 || reviewRating > 5){
        failed = "Rating must be in range from 1 to 5.";
        return;
    }else{
        failed = "";
    }
    if(reviewComment == "" || reviewComment == undefined ){
        failed = "Comment can't be empty.";
        return;
    }else{
        failed = "";
    }
    axios.post("http://127.0.0.1:8000/api/review", {
            eventId: Number(data.data),
            userId: userId,
            rating: reviewRating,
            comment: reviewComment
        }, config)
        .then(function (response) {
            console.log(response);
            success = response.data
        })
        .catch(function (error) {
            console.log(error);
            failed = error.response.data.error;
        });
}
</script>

<main>
    {#if res != "" && res != undefined}
        <img src={res.eventBanner} alt="" loading="lazy" class="banner">
        <img src={res.eventImage} alt="" loading="lazy" class="cover">
        <div>{res.eventName}</div>
        <div>{res.eventDate}</div>
        <div>{res.eventLocation}</div>
        <div>{res.eventDescription}</div>
        <div>Vietu skaits: {res.eventCapacity}</div>
    {/if}
    {#if allReviews != "" && allReviews != undefined}
        {#each allReviews as review}
            <div>
                <p>{review.UserName}</p>
                <p>{review.Rating}</p>
                <p>{review.Comment}</p>
            </div>
        {/each}
    {/if}
    <input type="text" bind:value={reviewComment}>
    <input type="number" name="" id="" min="1" max="5" step="1" bind:value={reviewRating}>
    <button on:click={postReview}>Post Review</button>
    {#if failed != undefined && failed != ""}
        <p class="error">{failed}</p>
    {/if}
    {#if success != undefined && success != ""}
        <p class="success">{success}</p>
    {/if}
</main>

<style>
.banner{
    max-width: 100%;
    border-radius: 0.2rem;
}
.cover{
    max-width: min(20rem, 100%);
    border-radius: 0.3rem;
}
</style>