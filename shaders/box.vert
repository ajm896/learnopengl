#version 330 core
layout (location = 0) in vec3 aPos;
layout (location = 1) in vec3 aNormal;

uniform mat4 model;
uniform mat4 view;
uniform mat4 proj;
uniform vec3 lightColor;
uniform vec3 viewPos;

out vec3 ourColor;
out vec3 Normal;
out vec3 FragPos;

void main()
{
    gl_Position = proj * view * model * vec4(aPos, 1.0);
    ourColor = lightColor;
    Normal = aNormal;
    FragPos = vec3(model * vec4(aPos, 1));
}