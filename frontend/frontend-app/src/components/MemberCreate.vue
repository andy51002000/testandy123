<template>
  <div>
    <h2>Create New Member</h2>
    <form @submit.prevent="createMember">
      <label for="username">Username:</label>
      <input type="text" id="username" v-model="username" required>
      <button type="submit">Create Member</button>
    </form>
    <p v-if="message">{{ message }}</p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      username: '',
      message: ''
    };
  },
  methods: {
    createMember() {
      // Prepare data to be sent in the request
      const requestData = {
        username: this.username
      };

      // Make API request to create a new member
      axios.post('https://edb0-104-198-4-130.ngrok-free.app/member', requestData)
        .then(response => {
          this.message = response.data.message;
        })
        .catch(error => {
          console.error('Error creating member:', error);
          this.message = 'Failed to create member';
        });
    }
  }
};
</script>

<style scoped>
/* Add styling as needed */
</style>
