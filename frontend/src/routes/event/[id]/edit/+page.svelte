<script>
export let data;
import axios from "$lib/axios";
import { onMount } from "svelte";
import "iconify-icon"
// export let data;
let res, failed, success, title, description, category, date, location, cover, banner, capacity, seatsRequired, allCategories;
onMount(async () => {
    try{
        //get event data
        const response = await axios.get(`/event/${data.data}`);
        res = response.data;
        title = res.eventName;
        description = res.eventDescription;
        category = res.eventCategory;
        date = res.eventDate;
        location = res.eventLocation;
        cover = res.eventImage;
        banner = res.eventBanner;
        capacity = res.eventCapacity;
        seatsRequired = Boolean(res.SeatsRequired);
        //get all categories
        const categories = await axios.get('/categories');
        allCategories = categories.data;
    }catch (e){
        console.error("Error:", e);
    }
})

function updateEventData(){
    axios.put(`/event/${data.data}`, {
            eventName: title,
            eventDescription: description,
            eventCategory: category,
            eventDate: date,
            eventLocation: location,
            eventImage: cover,
            eventBanner: banner,
            eventCapacity: capacity,
            seatsRequired: Number(seatsRequired)
        })
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

<main class="flex cen col">
    {#if res != "" && res != undefined}
    <div class="margin-t flex cen col">
        <label for="">
        Event name:
        <input type="text" name="" id="" bind:value={title}>
        </label>
        <label for="">
            Event description:
            <textarea name="" id="" cols="30" rows="10">{description}</textarea>
        </label>
        <label for="">
        Event category:
        <select name="" id=""bind:value={category}>
        <option value={res.eventCategory}>{res.eventCategory}</option>
        {#if allCategories != "" && allCategories != undefined}                
            {#each allCategories as cat}
                <option value={cat.Id}>{cat.CategoryName}</option>
            {/each}
        {/if}
        </select>
        </label>
        <label for="" >
            Event date:
            <input type="datetime-local" name="" id="" bind:value={date}>
        </label>
        <label for="">
            Event location:
            <input type="text" name="" id="" bind:value={location}>
        </label>
        <label for="">
            Event image/cover:
            <input type="text" name="" id="" bind:value={cover}>
        </label>
        <label for="">
            Event banner:
            <input type="text" name="" id="" bind:value={banner}>
        </label>
        <label for="">
            Are seats required for this event?
            <input type="checkbox" name="" id="" bind:checked={seatsRequired}>
        </label>
        <button on:click={updateEventData}>Update Event</button>
    </div>
    {/if}
    {#if success != "" && success != undefined}
        <p class="success">{success}</p>
    {/if}
    {#if failed != "" && failed != undefined}
        <p class="error">{failed}</p>
    {/if}
</main>

<style>
    .margin-t{
        margin-top: 5rem;
    }
    input, select, button{
        margin: 1rem 1rem;
        padding: 1rem;
        border: none;
        border-radius: 0.4rem;
        background-color: var(--bg);
    }
    textarea{
        border: none;
        border-radius: 0.4rem;
        background-color: var(--bg);
        resize: none;
    }
/*    input[type=checkbox]{
        width: 10rem;
        padding: 1rem;
        border: none;
        border-radius: 0.4rem;
        background-color: var(--bg);
    }*/
</style>