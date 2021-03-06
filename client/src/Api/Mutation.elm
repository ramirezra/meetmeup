-- Do not manually edit this file, it was auto-generated by dillonkearns/elm-graphql
-- https://github.com/dillonkearns/elm-graphql


module Api.Mutation exposing (..)

import Api.InputObject
import Api.Interface
import Api.Object
import Api.Scalar
import Api.ScalarCodecs
import Api.Union
import Graphql.Internal.Builder.Argument as Argument exposing (Argument)
import Graphql.Internal.Builder.Object as Object
import Graphql.Internal.Encode as Encode exposing (Value)
import Graphql.Operation exposing (RootMutation, RootQuery, RootSubscription)
import Graphql.OptionalArgument exposing (OptionalArgument(..))
import Graphql.SelectionSet exposing (SelectionSet)
import Json.Decode as Decode exposing (Decoder)


type alias RegisterRequiredArguments =
    { input : Api.InputObject.RegisterInput }


{-|

  - input -

-}
register : RegisterRequiredArguments -> SelectionSet decodesTo Api.Object.AuthResponse -> SelectionSet decodesTo RootMutation
register requiredArgs object_ =
    Object.selectionForCompositeField "register" [ Argument.required "input" requiredArgs.input Api.InputObject.encodeRegisterInput ] object_ identity


type alias LoginRequiredArguments =
    { input : Api.InputObject.LoginInput }


{-|

  - input -

-}
login : LoginRequiredArguments -> SelectionSet decodesTo Api.Object.AuthResponse -> SelectionSet decodesTo RootMutation
login requiredArgs object_ =
    Object.selectionForCompositeField "login" [ Argument.required "input" requiredArgs.input Api.InputObject.encodeLoginInput ] object_ identity


type alias CreateMeetupRequiredArguments =
    { input : Api.InputObject.NewMeetup }


{-|

  - input -

-}
createMeetup : CreateMeetupRequiredArguments -> SelectionSet decodesTo Api.Object.Meetup -> SelectionSet decodesTo RootMutation
createMeetup requiredArgs object_ =
    Object.selectionForCompositeField "createMeetup" [ Argument.required "input" requiredArgs.input Api.InputObject.encodeNewMeetup ] object_ identity


type alias UpdateMeetupRequiredArguments =
    { id : Api.ScalarCodecs.Id
    , input : Api.InputObject.UpdateMeetup
    }


{-|

  - id -
  - input -

-}
updateMeetup : UpdateMeetupRequiredArguments -> SelectionSet decodesTo Api.Object.Meetup -> SelectionSet decodesTo RootMutation
updateMeetup requiredArgs object_ =
    Object.selectionForCompositeField "updateMeetup" [ Argument.required "id" requiredArgs.id (Api.ScalarCodecs.codecs |> Api.Scalar.unwrapEncoder .codecId), Argument.required "input" requiredArgs.input Api.InputObject.encodeUpdateMeetup ] object_ identity


type alias DeleteMeetupRequiredArguments =
    { id : Api.ScalarCodecs.Id }


{-|

  - id -

-}
deleteMeetup : DeleteMeetupRequiredArguments -> SelectionSet Bool RootMutation
deleteMeetup requiredArgs =
    Object.selectionForField "Bool" "deleteMeetup" [ Argument.required "id" requiredArgs.id (Api.ScalarCodecs.codecs |> Api.Scalar.unwrapEncoder .codecId) ] Decode.bool
