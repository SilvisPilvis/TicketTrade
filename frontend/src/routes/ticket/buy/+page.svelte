<script>
	// import Stripe from "stripe";
	// export const stripe = new Stripe(import.meta.env.VITE_STRIPE_API_KEY);
	import { onMount } from "svelte";
	import {loadStripe} from '@stripe/stripe-js';
	import Cookies from "js-cookie";
	import axios from "axios";
	let stripe, success;
	onMount(async () => {
		// const stripe = await loadStripe(import.meta.env.VITE_STRIPE_API_KEY);
		stripe = await loadStripe(import.meta.env.VITE_STRIPE_API_KEY);
		checkout();
	})
	
	
	// let cardNumber, expiry, cvc, email;
	function checkout(){
		// let value = {
		// 	email: email
		// }
		// console.log(value);
		const config = {
            headers: { Authorization: `Bearer ${Cookies.get('token')}` }
        };
		axios.post('http://localhost:8000/checkout', {}, config)
        .then(function (response) {
            console.log(response);
            success = response.data;
			console.log(success);
			const id = success.Id;
			stripe.redirectToCheckout({sessionId: id});
        })
        .catch(function (error) {
            console.log(error);
            failed = error.response.data.error;
        });
		// fetch("http://localhost:8000/checkout", {
		// 	method: "POST",
		// }).then(
		// 	async response => {
		// 		const res = await response.json();
		// 		console.log(res);
		// 		const id = res.Id;
		// 		stripe.redirectToCheckout({sessionId: id});
		// 	}
		// )
	}
</script>

<!-- <main class="flex col cen">
	Varbut onMount checkout bet uz /checkout/:id
	<div class="margin-t">
		Email:
		<input type="text" bind:value={email}>
		<button on:click={checkout}>Checkout</button>
	</div>
</main>

<style>
    .margin-t{
        margin-top: 5rem;
    }
</style> -->