#version 420

layout(triangles) in;
layout (triangle_strip, max_vertices=9) out;

in VertexData {
    vec2 textureCoords;
    vec3 normal;
} VertexIn[];

out VertexData {
    vec2 textureCoords;
    vec3 normal;
} VertexOut;

void main()
{
    for(int i = 0; i < gl_in.length(); i++)
    {
        // copy attributes
        gl_Position = gl_in[i].gl_Position;
        VertexOut.normal = VertexIn[i].normal;
        VertexOut.textureCoords = VertexIn[i].textureCoords;

        // done with the vertex
        EmitVertex();
    }
    EndPrimitive();

// uncomment to make everything render 3 times
//
//    for(int i = 0; i < gl_in.length(); i++)
//    {
//        // copy attributes
//        gl_Position = gl_in[i].gl_Position + vec4(4.0, 0.0, 0.0, 0.0);
//        VertexOut.normal = VertexIn[i].normal;
//        VertexOut.textureCoords = VertexIn[i].textureCoords;
//
//        // done with the vertex
//        EmitVertex();
//    }
//    EndPrimitive();
//
//    for(int i = 0; i < gl_in.length(); i++)
//    {
//        // copy attributes
//        gl_Position = gl_in[i].gl_Position + vec4(-4.0, 0.0, 0.0, 0.0);
//        VertexOut.normal = VertexIn[i].normal;
//        VertexOut.textureCoords = VertexIn[i].textureCoords;
//
//        // done with the vertex
//        EmitVertex();
//    }
//    EndPrimitive();
}