-- Do not manually edit this file, it was auto-generated by dillonkearns/elm-graphql
-- https://github.com/dillonkearns/elm-graphql


module Api.InputObject exposing (..)

import Api.Interface
import Api.Object
import Api.Scalar
import Api.ScalarCodecs
import Api.Union
import Graphql.Internal.Builder.Argument as Argument exposing (Argument)
import Graphql.Internal.Builder.Object as Object
import Graphql.Internal.Encode as Encode exposing (Value)
import Graphql.OptionalArgument exposing (OptionalArgument(..))
import Graphql.SelectionSet exposing (SelectionSet)
import Json.Decode as Decode


buildLoginInput : LoginInputRequiredFields -> LoginInput
buildLoginInput required =
    { email = required.email, password = required.password }


type alias LoginInputRequiredFields =
    { email : String
    , password : String
    }


{-| Type for the LoginInput input object.
-}
type alias LoginInput =
    { email : String
    , password : String
    }


{-| Encode a LoginInput into a value that can be used as an argument.
-}
encodeLoginInput : LoginInput -> Value
encodeLoginInput input =
    Encode.maybeObject
        [ ( "email", Encode.string input.email |> Just ), ( "password", Encode.string input.password |> Just ) ]


buildMeetupFilter : (MeetupFilterOptionalFields -> MeetupFilterOptionalFields) -> MeetupFilter
buildMeetupFilter fillOptionals =
    let
        optionals =
            fillOptionals
                { name = Absent }
    in
    { name = optionals.name }


type alias MeetupFilterOptionalFields =
    { name : OptionalArgument String }


{-| Type for the MeetupFilter input object.
-}
type alias MeetupFilter =
    { name : OptionalArgument String }


{-| Encode a MeetupFilter into a value that can be used as an argument.
-}
encodeMeetupFilter : MeetupFilter -> Value
encodeMeetupFilter input =
    Encode.maybeObject
        [ ( "name", Encode.string |> Encode.optional input.name ) ]


buildNewMeetup : NewMeetupRequiredFields -> NewMeetup
buildNewMeetup required =
    { name = required.name, description = required.description }


type alias NewMeetupRequiredFields =
    { name : String
    , description : String
    }


{-| Type for the NewMeetup input object.
-}
type alias NewMeetup =
    { name : String
    , description : String
    }


{-| Encode a NewMeetup into a value that can be used as an argument.
-}
encodeNewMeetup : NewMeetup -> Value
encodeNewMeetup input =
    Encode.maybeObject
        [ ( "name", Encode.string input.name |> Just ), ( "description", Encode.string input.description |> Just ) ]


buildRegisterInput : RegisterInputRequiredFields -> RegisterInput
buildRegisterInput required =
    { username = required.username, email = required.email, password = required.password, confirmPassword = required.confirmPassword, firstName = required.firstName, lastName = required.lastName }


type alias RegisterInputRequiredFields =
    { username : String
    , email : String
    , password : String
    , confirmPassword : String
    , firstName : String
    , lastName : String
    }


{-| Type for the RegisterInput input object.
-}
type alias RegisterInput =
    { username : String
    , email : String
    , password : String
    , confirmPassword : String
    , firstName : String
    , lastName : String
    }


{-| Encode a RegisterInput into a value that can be used as an argument.
-}
encodeRegisterInput : RegisterInput -> Value
encodeRegisterInput input =
    Encode.maybeObject
        [ ( "username", Encode.string input.username |> Just ), ( "email", Encode.string input.email |> Just ), ( "password", Encode.string input.password |> Just ), ( "confirmPassword", Encode.string input.confirmPassword |> Just ), ( "firstName", Encode.string input.firstName |> Just ), ( "lastName", Encode.string input.lastName |> Just ) ]


buildUpdateMeetup : (UpdateMeetupOptionalFields -> UpdateMeetupOptionalFields) -> UpdateMeetup
buildUpdateMeetup fillOptionals =
    let
        optionals =
            fillOptionals
                { name = Absent, description = Absent }
    in
    { name = optionals.name, description = optionals.description }


type alias UpdateMeetupOptionalFields =
    { name : OptionalArgument String
    , description : OptionalArgument String
    }


{-| Type for the UpdateMeetup input object.
-}
type alias UpdateMeetup =
    { name : OptionalArgument String
    , description : OptionalArgument String
    }


{-| Encode a UpdateMeetup into a value that can be used as an argument.
-}
encodeUpdateMeetup : UpdateMeetup -> Value
encodeUpdateMeetup input =
    Encode.maybeObject
        [ ( "name", Encode.string |> Encode.optional input.name ), ( "description", Encode.string |> Encode.optional input.description ) ]