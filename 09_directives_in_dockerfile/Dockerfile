# ################### FROM ####################
# The FROM instruction specifies the base image that the container will be built on top of. 
FROM ubuntu:18.04

LABEL maintainer="Aleksey Kravchuk"

# ################## WORKDIR ##################
# In a Dockerfile, the WORKDIR instruction sets the working directory for any command that follows it in the Dockerfile.
# This means that any commands that are run in the container will be executed relative to the specified directory. 
WORKDIR /etc/nginx/

# #################### RUN ####################
RUN apt-get update && \
    apt-get install -y -q curl gnupg2
RUN curl http://nginx.org/keys/nginx_signing.key | apt-key add -

RUN apt-get update && \
    apt-get install -y -q nginx

# Always run a container with a non-root user.
# Running as root inside the container is running as root in the Docker host.
# If an attacker gains access to your container, they have access to all the root privileges and can perform several attacks against the Docker host.
# RUN addgroup --system nginx_app && \
#     adduser --system rbruser --group app

# #################### COPY ###################
# Use the COPY instruction to copy local files from the HOST machine to the IMAGE (current working directory).
# I.e.copy "nginx.conf" from directory with docker-project from the HOST machine to docker image
# COPY --chown=rbruser:nginx_app nginx.conf /etc/nginx/nginx.conf
COPY nginx.conf /etc/nginx/nginx.conf

# Or using WORKDIR:
# COPY nginx.conf .

# If you want to copy all the files from the host’s current directory to the container’s current directory, you can use the below command:
# COPY . .

# ########### ENTRYPOINT and CMD #############
# In Dockerfiles, both the "ENTRYPOINT" and "CMD" instruction are used to specify commands the container is supposed to run when starting up.
# So  whether we end our Dockerfile with 'CMD ["some_command.sh"]' OR 'ENRYPOINT ["some_command.sh"]' it produces the exact same result.
# But what’s the difference then? The first thing to know about these two instructions is that they can be used both  in the same Dockerfile.
# If you put multiple CMD instructions or multiple ENTRYPOINT only the last one will take effect.

# But if we use a CMD in combination with an ENTRYPOINT they will both  be taken into account when running the container.
# When we introduce an ENTRYPOINT, docker will first look at the command defined by the ENTRYPOINT and then use whatever we have in CMD as additional options for that command.
# So if in Dockerfile we have:
#                              ENTRYPOINT ["entrypoint_command_1.sh"]
#                              CMD ["PARAM_1", "PARAM_2"]
# Then command “docker run my_container_name” in this case will result in the execution: "entrypoint_command_1.sh PARAM_1 PARAM_2"

ENTRYPOINT ["nginx"]
CMD ["-g", "daemon off;"]

# ################## EXPOSE ##################
# Use the EXPOSE command to tell Docker which ports the container will listen on at runtime
EXPOSE 8901

# ################## VOLUME ##################
VOLUME /var/lib/nginx
