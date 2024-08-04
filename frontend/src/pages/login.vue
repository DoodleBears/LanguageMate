<template>
  <v-layout>
    <v-main height="900">
      <v-container class="h-100 pa-0" fluid>
        <v-row align="center" no-gutters>
          <v-col class="d-flex justify-center my-12" cols="12" lg="5">
            <v-responsive class="flex-1-1 pa-4" max-width="420">
              <v-img
                class="mx-auto mb-4"
                max-width="60"
                src="https://vuetifyjs.b-cdn.net/docs/images/logos/v.svg"
              />

              <h1 class="text-h4 text-center font-weight-bold mb-2">
                Welcome back!
              </h1>

              <div class="text-h6 font-weight-regular text-center mb-8">
                Log into your account
              </div>

              <div>
                <v-label>Email</v-label>

                <v-text-field
                  color="primary"
                  density="comfortable"
                  placeholder="Email"
                  variant="outlined"
                />
              </div>

              <div class="mb-8">
                <v-label>Password</v-label>

                <v-text-field
                  append-inner-icon="mdi-eye"
                  color="primary"
                  density="comfortable"
                  placeholder="Password"
                  variant="outlined"
                  @click:append-inner=""
                />

                <div class="text-end">
                  <v-btn
                    class="text-none me-n2"
                    color="primary"
                    slim
                    text="Forgot password?"
                    variant="text"
                  />
                </div>
              </div>

              <v-btn
                block
                class="text-none"
                color="primary"
                flat
                height="56"
                rounded="lg"
                text="Log In"
              />

              <v-divider class="my-12">
                <!-- remove in v3.6.7 -->
                <div class="text-no-wrap">Or continue with</div>
              </v-divider>

              <div class="d-flex align-center justify-center ga-4 mb-8">
                <v-btn color="#d93f21" icon="mdi-google" variant="tonal" />

                <v-btn color="#1867c0" icon="mdi-facebook" variant="tonal" />

                <v-btn icon="mdi-github" variant="tonal" />
              </div>

              <div class="text-center">
                Don't have an account?
                <v-btn
                  class="text-none"
                  color="primary"
                  slim
                  text="Sign up"
                  variant="text"
                />
              </div>
            </v-responsive>
          </v-col>

          <v-img
            class="flex-1-1 hidden-md-and-down"
            cover
            height="900"
            src="https://images.unsplash.com/photo-1596003903067-bf5762ad5c19?q=80&w=1470&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
          />
        </v-row>
      </v-container>
    </v-main>
  </v-layout>
</template>
