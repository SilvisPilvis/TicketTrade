// import stripe from "./stripe";
// export const POST = async ({request}) => {
// 	const data = await request.json();
// 	const cartItems = data.items;
// 	const lineItems = cartItems.map((item) => {
// 		return {
// 			price_data: {
// 				currency: "EUR",
// 				product_data:{
// 					name: item.name,
// 				},
// 				unit_amount: item.price * 100,
// 			},
// 			quantity: item.amount
// 		};
// 	})

// 	const session = await stripe.checkout.sessions.create({
// 		line_items: lineItems,
// 		shipping_address_collection: {
// 			allowed_countries: ["NO"],
// 		},
// 		mode: "payment",
// 		success_url: "/success",
// 		cancel_url: "/cancel",
// 		phone_number_collection: {
// 			enabled: true,
// 		},
// 	});

// 	return new Response(JSON.stringify({url: session.url}), {status: 200, headers: {
// 		"Conetnt-Type": "application/json"
// 	}});
// }