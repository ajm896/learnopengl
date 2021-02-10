#version 330 core
out vec4 FragColor;
  
in vec3 ourColor;
in vec2 TexCoord;

uniform sampler2D ourTexture;

void main()
{
    FragColor = vec4(1,0.01,1.0,1.0) * vec4(ourColor, 1);
}