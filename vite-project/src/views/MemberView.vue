<script setup>
    import SignInView from '../components/SignInView.vue'
    import SignUpView from '../components/SignUpView.vue'
    import UserDataView from '../components/UserDataView.vue'
</script>

<template>
    <div>

    <div v-if="isLogin === 'SignIn'">
        <SignInView @update="gotoRegister"/>
    </div>
    <div v-else-if="isLogin === 'ShowData'">
        <UserDataView :username="username" :account="account" :phone="phone"/>
    </div>
    <div v-else>
        <SignUpView />
    </div>
</div>
  
</template>


<script>
export default {
    components: {
        SignInView,
        SignUpView,
        UserDataView
    },
    data() {
        return {
            isLogin : "SignIn",
            username: "",
            account: "",
            phone: "",
        };
    },
    methods: {
        gotoRegister(newStatus, u){
            this.isLogin = newStatus
            console.log(newStatus, u)
            this.username = u.UserName
            this.account =  u.Account
            this.phone = u.Phone
        }
    }
};
</script>

<style>

</style>