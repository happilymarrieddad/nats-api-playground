<script lang="ts">
	import { defineComponent, ref } from 'vue';
	//import { useStore } from 'vuex';
	import { useRouter } from 'vue-router';
    import { useAuthStore } from '@/stores/auth'

	export default defineComponent({
		setup() {
			// const store = useStore();
            const auth = useAuthStore();
			const router = useRouter();

			const email = ref('');
			const password = ref('');
			const errMsg = ref('');

			const loginFn = () => {
				errMsg.value = '';

                auth.login(email.value, password.value);

                router.push({ name: 'home' });
			};

			return {
				email,
				password,
				errMsg,
				loginFn,
			};
		},
	});
</script>

<template>
	<form class="form-horizontal row g-3">
		<div class="col-md-12">
			<label for="exampleInputEmail1" class="form-label">Username</label>
			<input
				type="email"
				class="form-control"
				id="exampleInputEmail1"
				placeholder="user@mail.com"
				v-model="email"
			/>
		</div>
		<div class="col-md-12">
			<label for="exampleInputPassword1" class="form-label">Password</label>
			<input
				type="password"
				class="form-control"
				id="exampleInputPassword1"
				v-model="password"
			/>
		</div>
		<span style="color: red">{{ errMsg }}</span>
		<div class="col-md-12">
			<button type="button" class="btn btn-primary" @click="loginFn">
				Login
			</button>
		</div>
	</form>
</template>

<style scoped>
	.form-horizontal {
		display: block;
		width: 50%;
		margin: 0 auto;
	}
</style>