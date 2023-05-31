<script lang="ts">
	import { defineComponent } from 'vue';

	export default defineComponent({
		setup() {
			return {

			};
		},
	});
</script>

<template>
	<nav class="navbar navbar-expand-lg navbar-dark bg-dark">
		<div class="container">
			<a class="navbar-brand" href="">App</a>
			<button
				class="navbar-toggler"
				type="button"
				data-bs-toggle="collapse"
				data-bs-target="#navbarTogglerDemo02"
				aria-controls="navbarTogglerDemo02"
				aria-expanded="false"
				aria-label="Toggle navigation"
			>
				<span class="navbar-toggler-icon"></span>
			</button>
			<div class="collapse navbar-collapse" id="navbarTogglerDemo02">
				<ul class="navbar-nav">
					<li class="nav-item">
						<router-link class="nav-link" to="/">Home</router-link>
					</li>
					<li class="nav-item">
						<router-link class="nav-link" to="/about">About</router-link>
					</li>
				</ul>
				<ul class="navbar-nav ms-auto">
					<li class="nav-item">
						<router-link class="nav-link" to="/logout">Logout</router-link>
					</li>
				</ul>
			</div>
		</div>
	</nav>
	<router-view />
</template>

<style scoped>

</style>
