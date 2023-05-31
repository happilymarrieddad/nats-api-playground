<template>
	<div class="home">
		<h1>This is an home page</h1>
	</div>
</template>

<script lang="ts">
	import { Options, Vue } from 'vue-class-component';

	@Options({
		components: {},
	})
	export default class HomeView extends Vue {}
</script>