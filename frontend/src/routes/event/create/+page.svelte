<script>
export let data;
import axios from "$lib/axios";
import { onMount } from "svelte";
import "iconify-icon"
// export let data;
let failed, success, title, description, category, date, location, cover, banner, capacity, seatsRequired, allCategories, allTypes, createTypes = {};
onMount(async () => {
    try{
        //get event data
        // const response = await axios.get(`/event/${data.data}`);
        // res = response.data;
        // title = res.EventName;
        // description = res.EventDescription;
        // category = res.EventCategory;
        // date = res.EventDate;
        // location = res.EventLocation;
        // cover = res.EventImage;
        // banner = res.EventBanner;
        // capacity = res.EventCapacity;
        // seatsRequired = Boolean(res.SeatsRequired);
        //get all categories
        const categories = await axios.get('/categories');
        allCategories = categories.data;
        const tickettypes = await axios.get('/ticket/types');
        allTypes = tickettypes.data;
        console.log(allTypes)
    }catch (e){
        console.error("Error:", e);
    }
})


function createEvent(){
    const sumValues = obj => Object.values(obj).reduce((a, b) => a + b, 0);
    if(title == "" || title == undefined){
        failed = "Event name can't be empty.";
        return;
    }
    else{
        failed = "";
    }

    if(description == "" || description == undefined){
        failed = "Description can't be empty.";
        return;
    }
    else{
        failed = "";
    }

    if(category == null || category == undefined){
        failed = "Category can't be empty.";
        return;
    }
    else{
        failed = "";
    }
    if(date == "" || date == undefined){
        failed = "Date can't be empty.";
        return;
    }
    else{
        failed = "";
    }
    if(Date.parse(date)-Date.parse(new Date()) < 0){
        failed = "Date can't be in past.";
        return;
    }
    else{
        failed = "";
    }
    if(location == "" || location == undefined){
        failed = "Location can't be empty.";
        return;
    }
    else{
        failed = "";
    }
    if(cover == "" || cover == undefined){
        failed = "Cover can't be empty.";
        return;
    }
    else{
        failed = "";
    }
    if(banner == "" || banner == undefined){
        failed = "Banner can't be empty.";
        return;
    }
    else{
        failed = "";
    }
    if(capacity <= 0 || capacity == undefined || capacity == null){
        failed = "Capcity must be greather than 0.";
    }
    else{
        failed = "";
    }
    // console.log(String(JSON.stringify(createTypes)));
    if(sumValues(createTypes) <= 0 || capacity <=0){
        failed = "Number of tickets must be greater than 0 and/or sum of all ticket types must be greater than 0.";
        return;
    }
    else{
        failed = "";
    }
    if(sumValues(createTypes) != capacity){
        console.error("Number of tickets must match the sum of all ticket types.")
        failed = "Number of tickets must match the sum of all ticket types.";
        return;
    }else{
        failed = "";
    }

    axios.post('/event/create', {
            eventName: title,
            eventDescription: description,
            eventCategory: category,
            eventDate: date,
            eventLocation: location,
            eventImage: cover,
            eventBanner: banner,
            eventCapacity: capacity,
            // EventTicketTypes: JSON.stringify(createTypes),
            ticketTypes: JSON.stringify(createTypes),
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

function addTicketType(id, amount, test){
    if(amount > 0){
        test[id] = Number(amount);
    }
    if(amount == 0){
        delete test[id];
    }
}
</script>

<main class="flex col">
    <div class="margin-t flex col cen">
        <label for="">
            Event name:
            <input type="text" name="" id="" bind:value={title}>
        </label>
        <label for="">
            Event description:
            <textarea name="" id="" cols="30" rows="10" bind:value={description}></textarea>
        </label>
        <label for="">
            Event category:
            <select name="" id=""bind:value={category}>
                {#if allCategories != "" && allCategories != undefined}                
                    {#each allCategories as cat}
                        <option value={cat.Id}>{cat.CategoryName}</option>
                    {/each}
                {/if}
            </select>
        </label>
        <div class="flex row cen">
            <label for="" class="half">
                Event date:
                <input type="datetime-local" name="" class="half-r" id="" bind:value={date}>
            </label>
            <label for="" class="half">
                Event location:
                <input type="text" name="" id="" class="half-l" bind:value={location}>
            </label>
        </div>
        <div class="flex row">
            <label for="" class="half">
                Event image/cover:
                <input type="text" name="" id="" class="half-r" bind:value={cover}>
            </label>
            <label for="" class="half">
                Event banner:
                <input type="text" name="" id="" class="half-l" bind:value={banner}>
            </label>
        </div>
        <label for="">
            Are seats required for this event?
            <input type="checkbox" class="flex check" name="" id="" bind:checked={seatsRequired}>
        </label>
        <label for="">
            Number of Tickets:
            <input type="number" name="" id="" min="1" step="1" bind:value={capacity}>
        </label>
        <label for="">
            Enter the how much of each ticket you want to print:
            Ticket Types:
            {#if allTypes != "" && allTypes != undefined}
                {#each allTypes as type}
                <label for="" class="flex col cen">
                    {type.typeName}
                    <input type="number" name="" id="" min="0" step="1" onchange="" on:change={(e) => addTicketType(type.Id, e.target.value, createTypes)}>
                </label>
                {/each}
            {/if}
        </label>
        {#if success != "" && success != undefined}
            <p class="success">{success}</p>
        {/if}
        {#if failed != "" && failed != undefined}
            <p class="error">{failed}</p>
        {/if}
        <button on:click={createEvent}>Create Event</button>
    </div>
</main>

<style>
    /* .left{
        margin-left: 5rem;
    } */
    .margin-t{
        margin-top: 5rem;
        /* width: 100%; */
    }
    .check{
        display: flex;
    }
    .half{
        width: 40vw;
    }
    .half-l{
        width: 38vw;
        /* margin: 0 0 0 2vw; */
    }
    .half-r{
        width: 38vw;
        /* margin: 0 2vw 0 0; */
    }
    label{
        width: 80vw;
    }
    textarea{
        resize: none;
    }
    input, select, textarea{
        width: 100%;
        gap: 1rem;
        padding: 1rem;
        border: none;
        border-radius: 0.4rem;
        background-color: var(--fg);
        margin-bottom: 1rem;
    }
    button{
        width: 80vw;
        margin: 1rem;
        padding: 1rem;
        border-radius: 0.4rem;
        border: none;
        background-color: var(--button-fill);
    }
</style>