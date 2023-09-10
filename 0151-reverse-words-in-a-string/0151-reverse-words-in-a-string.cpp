class Solution {
public:
    string reverseWords(string s) {
        int n=s.size();
    
        string ans="";
        string temp="";
        for(int i=0;i<n;i++){
            if(s[i]==' '&&temp==""){
                continue;
            }
            else if(s[i]==' '&&temp!=""){
                cout<<temp<<endl;
                if(ans=="")ans=temp+ans;
                else ans=temp+" "+ans;
                temp="";
            }
            else if(s[i]!=' '){
                temp+=s[i];
            }
        }
       if(temp=="")return ans;
        else if(temp!=""&&ans=="") return temp+ans;
        return temp+" "+ans;

    }
};